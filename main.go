package main

import (
	"fmt"
	"restip/https"
	"restip/todo"
)

func main() {
	todo := todo.NewList()
	httpHandlers := https.NewHTTPHandlers(todo)
	httpSerer := https.NewHTTPServer(httpHandlers)

	if err := httpSerer.StartServer(); err != nil {
		fmt.Println("Failed", err)
	}
}
