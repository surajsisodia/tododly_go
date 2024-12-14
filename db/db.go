package db

import (
	"fmt"
	"log"
	"tododly/models"
	"tododly/utils"

	oracle "github.com/godoes/gorm-oracle"
	"gorm.io/gorm"
)

var Connections *gorm.DB

func GetSqlDbConnection() *gorm.DB {
	fmt.Println("Connecting to DB...")
	connectionString := "oracle://" + utils.DB_USERNAME + ":" + utils.DB_PASSWORD + "@" + utils.DB_HOST + ":" + utils.DB_PORT + "/" + utils.DB_SID

	db, err := gorm.Open(oracle.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error is sql.Open : %s", err)
	}

	db.Statement.RaiseErrorOnNotFound = true
	db.AutoMigrate(&models.User{}, &models.Task{}, &models.UserCredential{})
	db.Debug()

	fmt.Println("DB Alive")
	return db
}
