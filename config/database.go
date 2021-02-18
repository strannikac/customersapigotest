package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetDB() (*gorm.DB) {
	const (
		dbHost     = "localhost"
		dbPort     = "5432"
		dbUser     = "postgres"
		dbPassword = "postgres"
		dbName   = "postgres"
	)

	db, err := gorm.Open("postgres", 
		"host=" + dbHost + " port=" + dbPort + " user=" + dbUser + " dbname=" + dbName + " sslmode=disable password=" + dbPassword)

	if err != nil {
		panic("failed to connect database")
	}
	
	return db
}