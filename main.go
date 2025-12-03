package main

import (
	"fmt"
	"restip/https"
	"restip/todo"
)

func main() {
	todo := todo.NewList()
	httpHandlers := https.NewHTTPHandlers(todo)
	httpServer := https.NewHTTPServer(httpHandlers)
	fmt.Println("Сервер запущен")

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("Failed", err)
	}
}
