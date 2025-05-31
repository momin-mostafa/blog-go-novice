package dbhandler

import (
	userModel "backend/user"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func GetDBPointer() *gorm.DB {
	return db
}

func InitalizeDB() *gorm.DB {
	connStr := "host=localhost port=5432 user=tamim password=tamim dbname=backend sslmode=disable"

	db, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		panic("GROM FAILED TO OPEN")
	}

	fmt.Println("Connected to PostgreSQL!")

	fmt.Printf("Auto Migrating User")

	err = db.AutoMigrate(&userModel.User{})
	if err != nil {
		panic("AUTO MIGRATION FAILED FOR USER")
	}

	return db
}
