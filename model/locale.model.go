package model

type Locale struct {
	Id uint `gorm:"primary_key;auto_increment"`
	LanguageId uint `gorm:"not null;unique_index:idx_language_name"`
	Language Language
	Name string `gorm:"type:varchar(75);not null;unique_index:idx_language_name"`
	Value string `gorm:"not null"`
}

type LocaleModel struct {
}

func (model LocaleModel) Get(languageId uint) (map[string]string) {
	var locales []Locale
	Db.Where("language_id = ?", languageId).Find(&locales)

	var items = make(map[string]string)

	for _, item := range locales {
		items[item.Name] = item.Value
    }

	return items
}

func (model LocaleModel) Init(languages []Language) {
	var items []Locale

	for _, lang := range languages {
        if lang.Iso == "et" {
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_CUSTOMER", Value: "Klient"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_CUSTOMERS", Value: "Kliendid"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_MALE", Value: "Mees"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_FEMALE", Value: "Naine"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_ADD", Value: "Lisama"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_EDIT", Value: "Muuda"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_SAVE", Value: "Salvesta"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_DELETE", Value: "Kustuta"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_FNAME", Value: "Eesnimi"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_LNAME", Value: "Perekonnanimi"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_EMAIL", Value: "E-post"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_BIRTHDATE", Value: "Sünnikuupäev"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_GENDER", Value: "Sugu"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_ADDRESS", Value: "Aadress"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_ACTIONS", Value: "Toimingud"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_SEARCH", Value: "Otsi"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_CLEAR", Value: "Selge"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "ERR_MIN_SEARCH_LENGTH", Value: "Otsingustringi pikkus peab olema vähemalt 3 tähemärki."})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_DATA_NOT_FOUND", Value: "Andmeid ei leitud."})
            items = append(items, Locale{LanguageId: lang.Id, Name: "ERR_EMPTY_SEARCH_DATA", Value: "Otsingu jaoks andmed puuduvad."})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_TO_CUSTOMERS", Value: "Tagasi klientide juurde"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_DATA_NOT_SAVED", Value: "Andmeid ei salvestatud."})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_DATA_SAVED", Value: "Andmeid on salvestatud."})
            items = append(items, Locale{LanguageId: lang.Id, Name: "ERR_404", Value: "Viga 404"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "TXT_SOMETHING_WRONG", Value: "Midagi läks valesti..."})
        } else if lang.Iso == "ru" {
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_CUSTOMER", Value: "Пользователь"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_CUSTOMERS", Value: "Пользователи"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_MALE", Value: "Мужчина"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_FEMALE", Value: "Женщина"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_ADD", Value: "Добавить"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_EDIT", Value: "Редактировать"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_SAVE", Value: "Сохранить"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_DELETE", Value: "Удалить"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_FNAME", Value: "Имя"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_LNAME", Value: "Фамилия"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_EMAIL", Value: "E-mail"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_BIRTHDATE", Value: "Дата рождения"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_GENDER", Value: "Пол"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_ADDRESS", Value: "Адрес"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_ACTIONS", Value: "Действия"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_SEARCH", Value: "Искать"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_CLEAR", Value: "Очистить"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "ERR_MIN_SEARCH_LENGTH", Value: "Длина строки для поиска должна быть не менее 3 символов."})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_DATA_NOT_FOUND", Value: "Данные не были найдены."})
            items = append(items, Locale{LanguageId: lang.Id, Name: "ERR_EMPTY_SEARCH_DATA", Value: "Нет данных для поиска."})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_TO_CUSTOMERS", Value: "Назад к пользователям"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_DATA_NOT_SAVED", Value: "Данные не были сохранены."})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_DATA_SAVED", Value: "Данные были сохранены."})
            items = append(items, Locale{LanguageId: lang.Id, Name: "ERR_404", Value: "Ошибка 404"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "TXT_SOMETHING_WRONG", Value: "Что-то пошло не так..."})
		} else {
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_CUSTOMER", Value: "Customer"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_CUSTOMERS", Value: "Customers"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_MALE", Value: "Male"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_FEMALE", Value: "Female"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_ADD", Value: "Add"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_EDIT", Value: "Edit"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_SAVE", Value: "Save"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_DELETE", Value: "Delete"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_FNAME", Value: "First name"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_LNAME", Value: "Last name"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_EMAIL", Value: "E-mail"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_BIRTHDATE", Value: "Birth date"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_GENDER", Value: "Gender"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_ADDRESS", Value: "Address"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_ACTIONS", Value: "Actions"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_SEARCH", Value: "Search"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_CLEAR", Value: "Clear"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "ERR_MIN_SEARCH_LENGTH", Value: "Length of search string must be 3 or more characters."})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_DATA_NOT_FOUND", Value: "Data were not found."})
            items = append(items, Locale{LanguageId: lang.Id, Name: "ERR_EMPTY_SEARCH_DATA", Value: "No data for search."})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_TO_CUSTOMERS", Value: "Back to customers"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_DATA_NOT_SAVED", Value: "Data were not saved."})
            items = append(items, Locale{LanguageId: lang.Id, Name: "STR_DATA_SAVED", Value: "Data were saved."})
            items = append(items, Locale{LanguageId: lang.Id, Name: "ERR_404", Value: "Error 404"})
            items = append(items, Locale{LanguageId: lang.Id, Name: "TXT_SOMETHING_WRONG", Value: "Something went wrong..."})
		}
    }

	for index := range items {
		Db.Create(&items[index])
	}
}