package database

import (
	"assignment-2/models"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host = "localhost"
	dbPort = 5432
	user = "postgres"
	password = "postgres"
	dbname = "assignment-2"
	db *gorm.DB
	err error
)

func StartDB() (*gorm.DB, error) {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, 
password, dbname, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database :", err)
	}

	db.Debug().AutoMigrate(models.Orders{}, models.Items{})

	return db, nil
}
