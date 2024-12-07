package db

import (
	"fmt"
	"tododly/models"
	"tododly/utils"

	oracle "github.com/godoes/gorm-oracle"
	"gorm.io/gorm"
)

var Connections *gorm.DB

func init() {
	Connections = GetSqlDbConnection()
}

func GetSqlDbConnection() *gorm.DB {

	connectionString := "oracle://" + "fusion" + ":" + "welcome1" + "@" + utils.DB_HOST + ":" + "1521" + "/" + "FREE"

	// db, err := sql.Open("oracle", connectionString)
	db, err := gorm.Open(oracle.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(fmt.Errorf("error in sql.Open: %w", err))
	}

	db.Statement.RaiseErrorOnNotFound = true

	// err = db.Ping()
	// if err != nil {
	// 	panic(fmt.Errorf("error pinging db: %w", err))
	// }

	db.AutoMigrate(&models.User{}, &models.Task{}, &models.UserCredential{})
	// db.AutoMigrate(&models.Task{})

	db.Debug()

	fmt.Println("DB Alive")
	return db
}
