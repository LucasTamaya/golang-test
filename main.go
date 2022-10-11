package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"errors"
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

func getSingleTodo(ctx *gin.Context) {
	// get the id in the url params
	id := ctx.Param("id")
	todo, err := getTodoById(id)

	// if we don't found the todo, return this message
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	// if we found the todo
	ctx.IndentedJSON(http.StatusOK, todo)
}

func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("Todo not found")
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

func updateTodo(ctx *gin.Context) {
	var newData todo

	// get the body of the req and append it to newData's pointer
	if err := ctx.BindJSON(&newData); err != nil {
		return
	}

	// get the id in the url
	id := ctx.Param("id")

	todo, err := getTodoById(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found :("})
		return
	}

	// if we found a pointer for the Completed props, we update the todo.Completed
	if &newData.Completed != nil {
		todo.Completed = newData.Completed
	}

	// if we found a pointer for the Item props, we update the todo.Item
	if &newData.Item != nil {
		todo.Item = newData.Item
	}

	// finally we return the updated todo
	ctx.IndentedJSON(http.StatusOK, todo)

}

func main() {
	// create the router
	router := gin.Default()

	// create some routes
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getSingleTodo)
	router.PATCH("/todos/:id", updateTodo)
	router.POST("/todos", addTodo)

	// run the app at localhost port 8080
	router.Run("localhost:8080")

}
