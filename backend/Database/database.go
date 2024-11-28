package database

import (
	"fmt"
	"log"
	"os"
	models "proyecto/Models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Init() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
		panic(err)
	}

	user := os.Getenv("DATABASE_USER")
	pass := os.Getenv("DATABASE_PASS")
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	name := os.Getenv("DATABASE_NAME")

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to MySQL server:", err)
		panic(err)
	}

	err = db.Exec("CREATE DATABASE IF NOT EXISTS " + name).Error
	if err != nil {
		log.Fatal("Error creating database:", err)
		panic(err)
	}

	sqlDB, _ := db.DB()
	sqlDB.Close()

	dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, name)
	Database, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
		panic(err)
	}

	createSchema(Database)
}

func createSchema(db *gorm.DB) {

	err := db.AutoMigrate(
		&models.Bus{},
		&models.Conductor{},
		&models.Ruta{},
		&models.Asignacion{},
	)

	if err != nil {
		log.Fatal("Error creating schema:", err)
		panic(err)
	}

	log.Println("Schema created successfully")
}
