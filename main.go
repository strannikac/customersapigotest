package main

import (
    "config"
    "model"
    "helper"
    "log"
	"time"
	"bytes"
    "html/template"
    "net/http"
    "encoding/json"
    "strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
    "github.com/go-playground/validator"
)

var db *gorm.DB
var validate *validator.Validate
var response = make(map[string]string)
const statusSuccess = "success"

var customerModel model.CustomerModel

var languages []model.Language
var currentLanguage model.Language
var languageModel model.LanguageModel

var localeModel model.LocaleModel
var locales map[string]string
var genderValues = make(map[int]string)

const perPage = 3
var customersUrl = ""

func main() {
    db = config.GetDB()
    model.Db = db

    migrateDb()
    initDataDb()

    config.SetConstants()

    getLangs()
    initLocalesDb()

    handleRequests()

    db.Close()
}

func getLangs() {
    languages, currentLanguage = languageModel.Get()
}

func handleRequests() {
    router := mux.NewRouter()

    cssHandler := http.FileServer(http.Dir(config.HTML_PATH + "css/"))
    imagesHandler := http.FileServer(http.Dir(config.HTML_PATH + "images/"))
    jsHandler := http.FileServer(http.Dir(config.HTML_PATH + "js/"))

    http.Handle("/css/", http.StripPrefix("/css/", cssHandler))
    http.Handle("/images/", http.StripPrefix("/images/", imagesHandler))
    http.Handle("/js/", http.StripPrefix("/js/", jsHandler))

    router.HandleFunc("/{lang}/error404", showError404).Methods("GET")
    router.HandleFunc("/{lang}/error404/", showError404).Methods("GET")

    router.HandleFunc("/", homePage).Methods("GET")
    router.HandleFunc("/{lang}", homePage).Methods("GET")
    router.HandleFunc("/{lang}/", homePage).Methods("GET")
    router.HandleFunc("/{lang}/customers", homePage).Methods("GET")
    router.HandleFunc("/{lang}/customers/", homePage).Methods("GET")

    router.HandleFunc("/{lang}/add", htmlCustomerForm).Methods("GET")
    router.HandleFunc("/{lang}/add/", htmlCustomerForm).Methods("GET")
    router.HandleFunc("/{lang}/customers/add", htmlCustomerForm).Methods("GET")
    router.HandleFunc("/{lang}/customers/add/", htmlCustomerForm).Methods("GET")

    router.HandleFunc("/{lang}/edit/{id}", htmlCustomerForm).Methods("GET")
    router.HandleFunc("/{lang}/edit/{id}/", htmlCustomerForm).Methods("GET")
    router.HandleFunc("/{lang}/customers/edit/{id}", htmlCustomerForm).Methods("GET")
    router.HandleFunc("/{lang}/customers/edit/{id}/", htmlCustomerForm).Methods("GET")

    router.HandleFunc("/{lang}/{id}", htmlCustomer).Methods("GET")
    router.HandleFunc("/{lang}/{id}/", htmlCustomer).Methods("GET")
    router.HandleFunc("/{lang}/customers/{id}", htmlCustomer).Methods("GET")
    router.HandleFunc("/{lang}/customers/{id}/", htmlCustomer).Methods("GET")

    router.HandleFunc("/{lang}", saveCustomer).Methods("POST")
    router.HandleFunc("/{lang}/", saveCustomer).Methods("POST")
    router.HandleFunc("/{lang}/customers", saveCustomer).Methods("POST")
    router.HandleFunc("/{lang}/customers/", saveCustomer).Methods("POST")

    router.HandleFunc("/{lang}/{id}", saveCustomer).Methods("PUT")
    router.HandleFunc("/{lang}/{id}/", saveCustomer).Methods("PUT")
    router.HandleFunc("/{lang}/customers/{id}", saveCustomer).Methods("PUT")
    router.HandleFunc("/{lang}/customers/{id}/", saveCustomer).Methods("PUT")

    router.HandleFunc("/{lang}/{id}", deleteCustomer).Methods("DELETE")
    router.HandleFunc("/{lang}/{id}/", deleteCustomer).Methods("DELETE")
    router.HandleFunc("/{lang}/customers/{id}", deleteCustomer).Methods("DELETE")
    router.HandleFunc("/{lang}/customers/{id}/", deleteCustomer).Methods("DELETE")

    http.Handle("/", router)
    log.Fatal(http.ListenAndServe(":1000", nil))
}

func redirectErrorPage(w http.ResponseWriter, r *http.Request){
    setLanguage(r)
    http.Redirect(w, r, "/" + currentLanguage.Iso + "/error404/", http.StatusSeeOther)
}

func showError404(w http.ResponseWriter, r *http.Request){
    setLanguage(r)
    
    tpl := template.Must(template.ParseFiles(config.TPL_PATH + "error404.html"))

    data := struct {
		StrHeader string
		TxtDescription string
    }{
		StrHeader: locales["ERR_404"],
		TxtDescription: locales["TXT_SOMETHING_WRONG"], 
	}

    var tplBuffer bytes.Buffer
    tpl.Execute(&tplBuffer, data)

    showIndexTemplate(w, tplBuffer.String())
}

func showIndexTemplate(w http.ResponseWriter, content string) {
    tpl := template.Must(template.ParseFiles(config.TPL_PATH + "index.html"))

    type TplLanguage struct {
        Class string
        Link string
        Iso string
        Id uint
        Name string
    }

    var tplLanguages []TplLanguage

    for _, lang := range languages {
        var cssClass = ""
        if lang.Id == currentLanguage.Id {
            cssClass = "sel"
        }
        tplLanguages = append(tplLanguages, TplLanguage{Class: cssClass, Link: config.HOME_URL + lang.Iso + "/", Iso: lang.Iso, Id: lang.Id, Name: lang.Name})
    }

    jsonLocale, _ := json.Marshal(locales)

    data := struct {
		MetaTitle string
		MetaDescription string
		MetaKeywords string
		ApiUrl string
		HomeUrl string
		Locale string
		LanguageId uint
		LanguageIso string
		LanguageName string
		Content template.HTML
		Languages []TplLanguage
	}{
		MetaTitle: "Home page",
		MetaDescription: "Home page description",
		MetaKeywords: "Home page keywords",
		ApiUrl: config.API_URL,
		HomeUrl: config.HOME_URL,
		Locale: string(jsonLocale),
		LanguageId: currentLanguage.Id,
		LanguageIso: currentLanguage.Iso,
		LanguageName: currentLanguage.Name,
		Content: template.HTML(content),
		Languages: tplLanguages, 
	}

    tpl.Execute(w, data)
}

func homePage(w http.ResponseWriter, r *http.Request){
    setLanguage(r)

    items, paginator := htmlCustomers(w, r)

    tpl := template.Must(template.ParseFiles(config.TPL_PATH + "customers.html"))

    data := struct {
		StrCustomers string
		StrAdd string
		StrSearch string
        StrClear string
		StrFname string
		StrLname string
		StrEmail string
		StrBirthDate string
		StrGender string
		StrAddress string
		StrActions string
        LinkAdd string
        Items template.HTML
        Paginator template.HTML
    }{
		StrCustomers: locales["STR_CUSTOMERS"],
		StrAdd: locales["STR_ADD"],
		StrSearch: locales["STR_SEARCH"],
		StrClear: locales["STR_CLEAR"],
		StrFname: locales["STR_FNAME"],
		StrLname: locales["STR_LNAME"],
		StrEmail: locales["STR_EMAIL"],
		StrBirthDate: locales["STR_BIRTHDATE"],
		StrGender: locales["STR_GENDER"],
		StrAddress: locales["STR_ADDRESS"],
		StrActions: locales["STR_ACTIONS"],
        LinkAdd: customersUrl + "add/", 
        Items: template.HTML(items), 
        Paginator: template.HTML(paginator), 
	}

    var tplBuffer bytes.Buffer
    tpl.Execute(&tplBuffer, data)

    showIndexTemplate(w, tplBuffer.String())
}

func htmlCustomers(w http.ResponseWriter, r *http.Request) (string, string) {
    var page = 1
    var sort = "last_name"
    var isDesc = false
    var fnameSearch = ""
    var lnameSearch = ""

    arr := r.URL.Query()["page"]
    if len(arr) > 0 && len(arr[0]) > 0 {
        var n = helper.CheckPositiveDigit(arr[0])
        if n > 0 {
            page = n
        }
    }

    arr = r.URL.Query()["sort"]
    if len(arr) > 0 && len(arr[0]) > 0 {
        if arr[0] == "first_name" || arr[0] == "email" || arr[0] == "birth_date" {
            sort = arr[0]
        }
    }

    arr = r.URL.Query()["desc"]
    if len(arr) > 0 && len(arr[0]) > 0 && arr[0] == "1" {
        isDesc = true
    }

    arr = r.URL.Query()["fname"]
    if len(arr) > 0 && len(arr[0]) > 1 {
        fnameSearch = arr[0]
    }

    arr = r.URL.Query()["lname"]
    if len(arr) > 0 && len(arr[0]) > 1 {
        lnameSearch = arr[0]
    }

    customers, pages := customerModel.Get(page, perPage, sort, isDesc, fnameSearch, lnameSearch)

    tpl := template.Must(template.ParseFiles(config.TPL_PATH + "customers-items.html"))

    var pagesArray []int

    if pages > 1 {
        for i := 1; i <= pages; i++ {
            pagesArray = append(pagesArray, i)
        }
    }

    type TplItem struct {
		StrFname string
		StrLname string
		StrEmail string
		StrBirthDate string
		StrGender string
		StrAddress string
        Id uint
        FirstName string
        LastName string
        Email string
        BirthDate string
        Gender string
        Address string
        StrEdit string
        StrDelete string
        LinkOne string
        LinkEdit string
    }

    var items []TplItem

    for _, item := range customers {
        strId := strconv.FormatUint(uint64(item.Id), 10)

        items = append(items, TplItem{
            StrFname: locales["STR_FNAME"],
            StrLname: locales["STR_LNAME"],
            StrEmail: locales["STR_EMAIL"],
            StrBirthDate: locales["STR_BIRTHDATE"],
            StrGender: locales["STR_GENDER"],
            StrAddress: locales["STR_ADDRESS"],
            Id: item.Id, 
            FirstName: item.FirstName, 
            LastName: item.LastName, 
            Email: item.Email, 
            BirthDate: item.BirthDate.Format("02.01.2006"), 
            Gender: genderValues[item.Gender], 
            Address: item.Address, 
            StrEdit: locales["STR_EDIT"], 
            StrDelete: locales["STR_DELETE"], 
            LinkOne: customersUrl + strId + "/", 
            LinkEdit: customersUrl + "edit/" + strId + "/"}) 
    }

    dataItems := struct {
		Items []TplItem
    }{
		Items: items, 
	}

    var tplBufferItems bytes.Buffer
    tpl.Execute(&tplBufferItems, dataItems)

    tpl = template.Must(template.ParseFiles(config.TPL_PATH + "paginator.html"))

    dataPages := struct {
		Pages []int
        CurrentPage int
    }{
		Pages: pagesArray, 
        CurrentPage: page, 
	}

    var tplBufferPages bytes.Buffer
    tpl.Execute(&tplBufferPages, dataPages)

    return tplBufferItems.String(), tplBufferPages.String()
}

func htmlCustomer(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    setLanguage(r)

    id := helper.CheckPositiveDigit(params["id"])
    if id > 0 {
        item := customerModel.GetById(uint(id))

        if(item.Id > 0) {
            tpl := template.Must(template.ParseFiles(config.TPL_PATH + "customer-one.html"))

            data := struct {
                StrCustomer string
                StrBackToCustomers string
                StrFname string
                StrLname string
                StrEmail string
                StrBirthDate string
                StrGender string
                StrAddress string
                Id uint
                FirstName string
                LastName string
                Email string
                BirthDate string
                Gender string
                Address string
            }{
                StrCustomer: locales["STR_CUSTOMER"],
                StrBackToCustomers: locales["STR_TO_CUSTOMERS"],
                StrFname: locales["STR_FNAME"],
                StrLname: locales["STR_LNAME"],
                StrEmail: locales["STR_EMAIL"],
                StrBirthDate: locales["STR_BIRTHDATE"],
                StrGender: locales["STR_GENDER"],
                StrAddress: locales["STR_ADDRESS"],
                Id: item.Id, 
                FirstName: item.FirstName, 
                LastName: item.LastName, 
                Email: item.Email, 
                BirthDate: item.BirthDate.Format("02.01.2006"), 
                Gender: genderValues[item.Gender], 
                Address: item.Address,
            }

            var tplBuffer bytes.Buffer
            tpl.Execute(&tplBuffer, data)

            showIndexTemplate(w, tplBuffer.String())
        } else {
            redirectErrorPage(w, r)
        }
    } else {
        redirectErrorPage(w, r)
    }
}

func htmlCustomerForm(w http.ResponseWriter, r *http.Request) {
    setLanguage(r)
    params := mux.Vars(r)

    tpl := template.Must(template.ParseFiles(config.TPL_PATH + "customer-form.html"))
    var item model.Customer

    var strAction string
    strAction = locales["STR_EDIT"]

    id := helper.CheckPositiveDigit(params["id"])
    if id > 0 {
        //edit
        item = customerModel.GetById(uint(id))

        if(item.Id > 0) {
        } else {
            //TODO: error 404 
            redirectErrorPage(w, r)
        }
    } else {
        //add
        strAction = locales["STR_ADD"]

        item = model.Customer{
            Id: 0, 
            FirstName: "", 
            LastName: "", 
            BirthDate: time.Now(), 
            Gender: 1, 
            Email: "", 
            Address: "", 
        }
    }

    data := struct {
        StrAction string
        StrSave string
        StrBackToCustomers string
        StrFname string
        StrLname string
        StrEmail string
        StrBirthDate string
        StrGender string
        StrAddress string
        StrMale string
        StrFemale string
        LinkBack string
        Id uint
        FirstName string
        LastName string
        Email string
        BirthDate string
        Gender int
        Address string
    }{
        StrAction: strAction,
        StrSave: locales["STR_SAVE"],
        StrBackToCustomers: locales["STR_TO_CUSTOMERS"],
        StrFname: locales["STR_FNAME"],
        StrLname: locales["STR_LNAME"],
        StrEmail: locales["STR_EMAIL"],
        StrBirthDate: locales["STR_BIRTHDATE"],
        StrGender: locales["STR_GENDER"],
        StrAddress: locales["STR_ADDRESS"],
        StrMale: locales["STR_MALE"],
        StrFemale: locales["STR_FEMALE"],
        LinkBack: customersUrl, 
        Id: item.Id, 
        FirstName: item.FirstName, 
        LastName: item.LastName, 
        Email: item.Email, 
        BirthDate: item.BirthDate.Format("02.01.2006"), 
        Gender: item.Gender, 
        Address: item.Address,
    }

    var tplBuffer bytes.Buffer
    tpl.Execute(&tplBuffer, data)

    showIndexTemplate(w, tplBuffer.String())
}

func saveCustomer(w http.ResponseWriter, r *http.Request) {
    setLanguage(r)
    r.ParseMultipartForm(0)

    strId := r.Form.Get("id")
    fname := r.Form.Get("fname")
    lname := r.Form.Get("lname")
    email := r.Form.Get("email")
    strBirthdate := r.Form.Get("birthdate")
    gender := helper.CheckPositiveDigit(r.Form.Get("gender"))
    address := r.Form.Get("address")

    var format = "02.01.2006"
    if currentLanguage.Iso == "en" {
        format = "01/02/2006"
    }

    birthdate, _ := time.Parse(format, strBirthdate)

    if gender != 2 {
        gender = 1
    }

    validate = validator.New()
    var item = model.Customer{FirstName: fname, LastName: lname, BirthDate: birthdate, Gender: gender, Email: email, Address: address}
    err := validate.Struct(item)
    response = make(map[string]string)

    if err == nil {
        id := helper.CheckPositiveDigit(strId)
        if id > 0 {
            //update
            current := customerModel.GetById(uint(id))

            if(current.Id > 0) {
                item.Id = current.Id

                count := customerModel.Update(item)

                if count > 0 {
                    response["status"] = statusSuccess 
                    response["msg"] = locales["STR_DATA_SAVED"] 
                } else {
                    response["error"] = locales["STR_DATA_NOT_SAVED"]
                }
            } else {
                response["error"] = locales["STR_DATA_NOT_SAVED"]
            }
        } else {
            //create
            newId := customerModel.Create(item)
            if newId > 0 {
                response["id"] = strconv.FormatUint(uint64(newId), 10)
                response["status"] = statusSuccess 
                response["msg"] = locales["STR_DATA_SAVED"] 
            } else {
                response["error"] = locales["STR_DATA_NOT_SAVED"]
            }
        }
    } else {
        response["error"] = locales["STR_DATA_NOT_SAVED"]
    }

    setResponse(w)
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id := helper.CheckPositiveDigit(params["id"])

    if id > 0 {
        count := customerModel.Delete(uint(id))

        if count > 0 {
            response["status"] = statusSuccess 
            response["msg"] = locales["STR_DATA_SAVED"] 
        } else {
            response["error"] = locales["STR_DATA_NOT_SAVED"]
        }
    } else {
        response["error"] = locales["STR_DATA_NOT_SAVED"]
    }

    setResponse(w)
}

func setResponse(w http.ResponseWriter) {
    currentTime := time.Now()
    response["time"] = currentTime.Format("02.01.2006 15:04:05")
    json.NewEncoder(w).Encode(&response)
}

func setLanguage(r *http.Request) {
    params := mux.Vars(r)
    iso := params["lang"]

    if iso != "" {
        var i = languageModel.Check(iso, languages)
        if i >= 0 {
            currentLanguage = languages[i]
        }
    }

    locales = localeModel.Get(currentLanguage.Id)

    genderValues[1] = locales["STR_MALE"]
    genderValues[2] = locales["STR_FEMALE"]

    customersUrl = "/" + currentLanguage.Iso + "/customers/"
}

func migrateDb() {
    db.AutoMigrate(&model.Customer{})
	db.AutoMigrate(&model.Language{})
	db.AutoMigrate(&model.Locale{})
}

func initDataDb() {
    //add languages (if needed)
    languageModel.Init()
    //add cusomers (if needed)
    customerModel.Init()
}

func initLocalesDb() {
    //add locales (if needed)
    localeModel.Init(languages)
}