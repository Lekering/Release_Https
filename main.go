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

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("Failed", err)
		return
	}
	fmt.Println("Сервер запущен")

	// Блокируем выполнение, чтобы программа не завершилась
	select {}
}
