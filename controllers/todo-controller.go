package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/x0000ff/todo-api/models"
)

// TodoController ...
type TodoController struct {
	DB *gorm.DB
}

// Create creates new todo
func (tc *TodoController) Create(c *gin.Context) {

	db := tc.DB
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

// Show fetch single todo
func (tc *TodoController) Show(c *gin.Context) {

	db := tc.DB

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

// Update updates single todo
func (tc *TodoController) Update(c *gin.Context) {

	db := tc.DB
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

// Delete deletes single todo
func (tc *TodoController) Delete(c *gin.Context) {

	db := tc.DB
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

// Index fetch all todos
func (tc *TodoController) Index(c *gin.Context) {

	db := tc.DB

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
