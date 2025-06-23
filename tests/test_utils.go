package tests

import (
    "fmt"
    "os"
    "hospitalapp/models"
    "path/filepath"
    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/gin-gonic/gin"
)

var TestDB *gorm.DB

func SetupTestDB() *gorm.DB {
     absPath, _ := filepath.Abs("../.env.test")
     err := godotenv.Load(absPath)

    if err != nil {
        panic("⚠️ Failed to load .env.test: " + err.Error())
    }
	fmt.Println("DB_USER:", os.Getenv("DB_USER"))

    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_PORT"),
    )
	fmt.Println(dsn)
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to test database")
    }

    db.Exec("DROP SCHEMA public CASCADE; CREATE SCHEMA public;")
    db.AutoMigrate(&models.User{}, &models.Patient{})

    TestDB = db
    return db
}

func SetupRouter() *gin.Engine {
    gin.SetMode(gin.TestMode)
    return gin.Default()
}
