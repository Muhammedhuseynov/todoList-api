package controllerAPI

import (
	"github.com/Muhammedhuseynov/ginAPItodo_app/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllTodoList(c *gin.Context) {
	var todo []models.Todo
	err := models.GetAllTodos(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func GetATodo(c *gin.Context) {
	id := c.Params.ByName("id")
	var todolist models.Todo
	err := models.GetATodo(&todolist, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todolist)
	}
}

func CreateTodoList(c *gin.Context) {
	var todolist models.Todo
	c.BindJSON(&todolist)
	err := models.CreateTodoList(&todolist)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusOK, todolist)
	}
}

func UpdateTodoList(c *gin.Context) {
	var todolist models.Todo
	id := c.Params.ByName("id")

	err := models.GetATodo(&todolist, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, todolist)
	}
	c.BindJSON(&todolist)
	err = models.UpdateTodoList(&todolist, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todolist)
	}
}

func DeleteTodoListItem(c *gin.Context) {
	var todolist models.Todo
	id := c.Params.ByName("id")
	err := models.DeleteToListItem(&todolist, id)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		c.JSON(http.StatusOK, gin.H{"id " + id: "Deleted!"})
	}
}
