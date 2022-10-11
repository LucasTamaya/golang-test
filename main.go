package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"` // to format to json when sending to the client
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Clean up the room", Completed: false},
	{ID: "2", Item: "Reading", Completed: true},
	{ID: "3", Item: "Sports", Completed: false},
}

func getTodos(ctx *gin.Context) {
	// return the todos with JSON format
	ctx.IndentedJSON(http.StatusOK, todos)
}

func addTodo(ctx *gin.Context) {
	// init an empty variable of type todo
	var newTodo todo

	// ctx.BindJSON get the body of the request and append it to the memory address of newTodo
	if err := ctx.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	ctx.IndentedJSON(http.StatusCreated, newTodo)
}

func main() {
	// create the router
	router := gin.Default()

	// create some routes
	router.GET("/todos", getTodos)
	router.GET("/todos")
	router.POST("/todos", addTodo)

	// run the app at localhost port 8080
	router.Run("localhost:8080")

}
