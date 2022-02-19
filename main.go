package main

import (
	"github.com/Muhammedhuseynov/todoList-api/configData"
	"github.com/Muhammedhuseynov/todoList-api/controllerAPI"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	var _, err = configData.ConnectDB()
	if err != nil {
		log.Println("Not connected!")
	}
	log.Println("connected!")

	router := gin.Default()

	g := router.Group("/api")
	{
		g.GET("todolist", controllerAPI.GetAllTodoList)
		g.GET("todolist/:id", controllerAPI.GetATodo)
		g.POST("todolist/create", controllerAPI.CreateTodoList)
		g.PUT("todolist/update/:id", controllerAPI.UpdateTodoList)
		g.DELETE("todolist/delete/:id", controllerAPI.DeleteTodoListItem)
	}
	router.Run()
}
