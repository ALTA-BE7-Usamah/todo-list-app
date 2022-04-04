package utils

import (
	"fmt"
	"project2/todo-list-app/configs"
	"project2/todo-list-app/entities"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(config *configs.AppConfig) *gorm.DB {

	err := godotenv.Load()
	if err != nil {
		panic("Tidak ada file konfigurasi dotenv tersedia")
	}

	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		config.Database.Username,
		config.Database.Password,
		config.Database.Address,
		config.Database.Port,
		config.Database.Name,
	)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Info("failed to connect database :", err)
		panic(err)
	}

	InitialMigration(db)
	return db
}

func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.Task{})
	db.AutoMigrate(&entities.Project{})
}
