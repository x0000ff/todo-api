package main

import (
	"net/http"
	"strconv"

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
	v1 := router.Group("/api/v1/todos")
	{
		v1.POST("/", createTodo)
		v1.GET("/", fetchAllTodos)
		v1.GET("/:id", fetchSingleTodo)
		v1.PUT("/:id", updateTodo)
		v1.DELETE("/:id", deleteTodo)
	}

	router.Run()
}

func createTodo(c *gin.Context) {

	completed, _ := strconv.Atoi(c.PostForm("completed"))

	todo := models.TodoModel{
		Title:     c.PostForm("title"),
		Completed: completed,
	}

	db.Create(&todo)
	// db.Save(&todo)

	c.JSON(http.StatusCreated,
		gin.H{
			"status":     http.StatusCreated,
			"message":    "Todo item created successfully!",
			"resourceId": todo.ID,
		})
}

// fetchSingleTodo fetch single todo
func fetchSingleTodo(c *gin.Context) {

	var todo models.TodoModel
	id := c.Param("id")

	db.First(&todo, id)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "No todo found",
		})
		return
	}

	completed := (todo.Completed == 1)

	apiTodo := models.ApiTodo{
		ID:        todo.ID,
		Title:     todo.Title,
		Completed: completed,
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   apiTodo,
	})
}

// updateTodo updates single todo
func updateTodo(c *gin.Context) {

	id := c.Param("id")

	var todo models.TodoModel
	db.First(&todo, id)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "No todo found",
		})
		return
	}

	db.Model(&todo).Update("title", c.PostForm("title"))

	completed, _ := strconv.ParseBool(c.PostForm("completed"))
	completedAsInt := 0
	if completed {
		completedAsInt = 1
	}
	db.Model(&todo).Update("completed", completedAsInt)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Todo updated succesfully",
	})
}

// deleteTodo deletes single todo
func deleteTodo(c *gin.Context) {

	id := c.Param("id")

	var todo models.TodoModel
	db.First(&todo, id)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Todo not found",
		})
		return
	}

	db.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Todo deleted succesfully",
	})
}

// fetchAllTodo fetch all todos
func fetchAllTodos(c *gin.Context) {

	var todos []models.TodoModel
	var apiTodos = []models.ApiTodo{}

	db.Find(&todos)

	// transforms the todos for building a good response
	for _, item := range todos {
		completed := (item.Completed == 1)
		apiTodos = append(apiTodos, models.ApiTodo{
			ID:        item.ID,
			Title:     item.Title,
			Completed: completed,
		})
	}

	c.JSON(http.StatusOK,
		gin.H{
			"status": http.StatusOK,
			"data":   apiTodos,
		})
}
