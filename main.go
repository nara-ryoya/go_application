package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Task struct {
	ID          uint      `gorm:"primary_key"`
	Task        string    `gorm:"size:255"`
	IsCompleted bool      `gorm:"default:false"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", dbHost, dbUser, dbPassword, dbName, dbPort)

	// wait for db to be ready
	time.Sleep(3 * time.Second)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	db.AutoMigrate(&Task{})

	log.Println("starting server..")
	r := gin.Default()
	r.GET("/hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "hello world!",
		})
	})
	r.GET("/tasks", func(c *gin.Context) {
		var tasks []Task
		db.Find(&tasks)
		c.JSON(http.StatusOK, tasks)
	})
	r.POST("/tasks", func(c *gin.Context) {
		var task Task
		if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&task)
		c.JSON(http.StatusCreated, task)
	})
	r.PUT("/tasks/:id", func(c *gin.Context) {
		var task Task
		id := c.Param("id")

		if err := db.First(&task, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
			return
		}

		if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db.Save(&task)
		c.JSON(http.StatusOK, task)
	})
	r.DELETE("/tasks/:id", func(c *gin.Context) {
		var task Task
		id := c.Param("id")

		if err := db.First(&task, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
			return
		}

		db.Delete(&task)
		c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully!"})
	})
	log.Fatal(r.Run(":8080"))
}
