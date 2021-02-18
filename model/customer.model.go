package model

import (
	"time"
    "math"
)

type Customer struct {
	Id uint `gorm:"primary_key;auto_increment;not null"`
	FirstName string `gorm:"type:varchar(100);not null" validate:"required,alpha"`
	LastName string `gorm:"type:varchar(100);not null" validate:"required,alpha"`
	BirthDate time.Time `validate:"required"`
	Gender int `gorm:"not null;default:1" validate:"required`
	Email string `gorm:"type:varchar(100);not null;unique" validate:"required,email"`
	Address string `gorm:"type:varchar(200)"`
}

type CustomerModel struct {
}

func (model CustomerModel) Get(
	page int, 
	perPage int, 
	sort string, 
	isDesc bool, 
	fname string, 
	lname string) ([]Customer, int) {

	var customers []Customer
	var count int = 0
	var sortby = "last_name"
	var limit = 10
	var offset = 0

	if perPage > 0 {
		limit = perPage
	}

	if page > 1 {
		offset = page * limit - limit
	}

	if sort == "first_name" || sort == "email" || sort == "birth_date" {
		sortby = sort
	}

	if isDesc {
		sortby = sortby + " DESC"
	}

	if lname != "" && fname != "" {
		Db.Model(&Customer{}).Where("last_name LIKE ? AND first_name LIKE ?", "%" + lname + "%", "%" + fname + "%").Count(&count)
		Db.Model(&Customer{}).Where("last_name LIKE ? AND first_name LIKE ?", "%" + lname + "%", "%" + fname + "%").Order(sortby).Limit(limit).Offset(offset).Find(&customers)
	} else if lname != "" {
		Db.Model(&Customer{}).Where("last_name LIKE ?", "%" + lname + "%").Count(&count)
		Db.Model(&Customer{}).Where("last_name LIKE ?", "%" + lname + "%").Order(sortby).Limit(limit).Offset(offset).Find(&customers)
	} else if fname != "" {
		Db.Model(&Customer{}).Where("first_name LIKE ?", "%" + fname + "%").Count(&count)
		Db.Model(&Customer{}).Where("first_name LIKE ?", "%" + fname + "%").Order(sortby).Limit(limit).Offset(offset).Find(&customers)
	} else {
		Db.Model(&Customer{}).Count(&count)
		Db.Model(&Customer{}).Order(sortby).Limit(limit).Offset(offset).Find(&customers)
	}

	pages := math.Ceil(float64(count) / float64(limit))

	return customers, int(pages)
}

func (model CustomerModel) GetById(id uint) (Customer) {
	var customer Customer

	Db.First(&customer, id)

	return customer
}

func (model CustomerModel) Create(customer Customer) (int) {
	Db.Create(&customer)

	return int(customer.Id)
}

func (model CustomerModel) Update(customer Customer) (int64) {
	result := Db.Save(&customer)

	return result.RowsAffected
}

func (model CustomerModel) Delete(id uint) (int64) {
	result := Db.Delete(&Customer{}, id)

	return result.RowsAffected
}

func (model CustomerModel) Init() {
	var format = "2006-01-02"
	var items []Customer

	date, _ := time.Parse(format, "1999-03-03")
	var item = Customer{FirstName: "John", LastName: "Smith", BirthDate: date, Gender: 1, Email: "john.smith@test.com", Address: ""}
	items = append(items, item)

	date, _ = time.Parse(format, "1995-05-05")
	item = Customer{FirstName: "Samantha", LastName: "Fox", BirthDate: date, Gender: 2, Email: "samantha@test.com", Address: ""}
	items = append(items, item)

	date, _ = time.Parse(format, "1989-09-09")
	item = Customer{FirstName: "Jim", LastName: "Scott", BirthDate: date, Gender: 1, Email: "jim.scott@test.com", Address: ""}
	items = append(items, item)

	date, _ = time.Parse(format, "1983-07-21")
	item = Customer{FirstName: "Иван", LastName: "Иванов", BirthDate: date, Gender: 1, Email: "ivan.ivanov@ya.ru", Address: ""}
	items = append(items, item)

	date, _ = time.Parse(format, "2001-11-11")
	item = Customer{FirstName: "Владимир", LastName: "Сидоров", BirthDate: date, Gender: 1, Email: "v.sidorov@mail.ru", Address: ""}
	items = append(items, item)

	date, _ = time.Parse(format, "2003-07-15")
	item = Customer{FirstName: "Samanta", LastName: "Lill", BirthDate: date, Gender: 2, Email: "samanta.lill@test.com", Address: ""}
	items = append(items, item)

	date, _ = time.Parse(format, "1991-01-30")
	item = Customer{FirstName: "Екатерина", LastName: "Иванова", BirthDate: date, Gender: 2, Email: "e.ivanova@ya.ru", Address: ""}
	items = append(items, item)

	date, _ = time.Parse(format, "2000-12-23")
	item = Customer{FirstName: "Alex", LastName: "Smart", BirthDate: date, Gender: 1, Email: "alex.smart@test.com", Address: ""}
	items = append(items, item)

	date, _ = time.Parse(format, "2002-12-11")
	item = Customer{FirstName: "Alexander", LastName: "Smith", BirthDate: date, Gender: 1, Email: "a.smith@test.com", Address: ""}
	items = append(items, item)

	date, _ = time.Parse(format, "1989-10-11")
	item = Customer{FirstName: "Alexey", LastName: "Aleksejev", BirthDate: date, Gender: 1, Email: "a.aleksejev@test.com", Address: ""}
	items = append(items, item)

	date, _ = time.Parse(format, "1999-10-11")
	item = Customer{FirstName: "Andreas", LastName: "Magelatti", BirthDate: date, Gender: 1, Email: "a.magelatti@test.com", Address: ""}
	items = append(items, item)

	date, _ = time.Parse(format, "1994-11-11")
	item = Customer{FirstName: "Toomas", LastName: "Pirk", BirthDate: date, Gender: 1, Email: "toomas.pirk@test.com", Address: ""}
	items = append(items, item)

	date, _ = time.Parse(format, "1996-09-17")
	item = Customer{FirstName: "Arvo", LastName: "Pikk", BirthDate: date, Gender: 1, Email: "arvo.pikk@test.com", Address: ""}
	items = append(items, item)

	date, _ = time.Parse(format, "1999-06-07")
	item = Customer{FirstName: "Anna", LastName: "Pikk", BirthDate: date, Gender: 2, Email: "anna.pikk@test.com", Address: ""}
	items = append(items, item)

	date, _ = time.Parse(format, "1999-06-05")
	item = Customer{FirstName: "Kristina", LastName: "Sidorova", BirthDate: date, Gender: 2, Email: "k.sidorova@test.com", Address: ""}
	items = append(items, item)

	date, _ = time.Parse(format, "2001-07-17")
	item = Customer{FirstName: "Konstantin", LastName: "Fanin", BirthDate: date, Gender: 1, Email: "k.fanin@test.com", Address: ""}
	items = append(items, item)

	date, _ = time.Parse(format, "2003-05-19")
	item = Customer{FirstName: "Ksenia", LastName: "Belousova", BirthDate: date, Gender: 2, Email: "k.belousova@test.com", Address: ""}
	items = append(items, item)

	date, _ = time.Parse(format, "1990-12-11")
	item = Customer{FirstName: "Alexander", LastName: "Johnson", BirthDate: date, Gender: 1, Email: "a.johnson@test.com", Address: ""}
	items = append(items, item)

	date, _ = time.Parse(format, "1992-08-11")
	item = Customer{FirstName: "Andrey", LastName: "Andreev", BirthDate: date, Gender: 1, Email: "a.andreev@test.com", Address: ""}
	items = append(items, item)

	for index := range items {
		Db.Create(&items[index])
	}
}