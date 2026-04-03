package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Hello, World!")
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", handler)     // регистрируем маршрут
	http.ListenAndServe(":8080", nil) // запускаем сервер на порту 8080
}
