package database

import (
	"fmt"
	"log"
	"time"

	"github.com/meajudaaqui/user-service/database/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	str := "host=localhost port=25432 user=postgres dbname=users sslmode=disable password=postgres"

	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})
	if err != nil {
		fmt.Println("nao foi possível conectar ao banco de dados")
		log.Fatal("Error: ", err)
	}

	db = database
	db = database
	config, _ := db.DB()
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	migrations.RunMigrations(db)

}

func CloseConn() error {
	config, err := db.DB()
	if err != nil {
		return err
	}

	err = config.Close()
	if err != nil {
		return err
	}

	return nil
}

func GetDatabase() *gorm.DB {
	return db
}
