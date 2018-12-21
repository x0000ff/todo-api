package main

import (
	"github.com/x0000ff/todo-api/controllers"
	"github.com/x0000ff/todo-api/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func init() {

	// open a db connection
	var err error
	db, err = gorm.Open("sqlite3", "./db.db")
	if err != nil {
		panic("failed to cottect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.TodoModel{})
}

func main() {

	defer db.Close()

	router := gin.Default()
	todoController := &controllers.TodoController{DB: db}

	v1 := router.Group("/api/v1/todos")
	{
		v1.POST("/", todoController.Create)
		v1.GET("/", todoController.Index)
		v1.GET("/:id", todoController.Show)
		v1.PUT("/:id", todoController.Update)
		v1.DELETE("/:id", todoController.Delete)
	}

	router.Run()
}
