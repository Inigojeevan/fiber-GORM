package database

import (
	"log"
	"os"

	"github.com/Inigojeevan/fiber-GORM/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB *gorm.DB
)

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to create database \n", err.Error())
		os.Exit(2)
	}
	log.Println("Connected to database")
	db.Logger = logger.Default.LogMode(logger.Info) //logger settings

	db.AutoMigrate(&models.Order{}, &models.Products{}, &models.User{})

	DB = db
}
