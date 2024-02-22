package config

import (
	"github.com/itsmaheshkariya/gin-gorm-rest/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
func Connect() {
    db, err := gorm.Open(postgres.Open("host=localhost port=5432 user=postgres password=0852138201 dbname=postgres sslmode=disable"), &gorm.Config{})
    if err != nil {
        panic("Could not connect to the database")
    }
    db.AutoMigrate(&model.User{})
    DB = db
}
