package main

import (
	"log"
	"net/http"
	"poll_service/service"
)

func main() {
	// Создаем структуру для нашего приложения
	// содержит gorm.DB с методами
	app := service.Service{}
	if err := app.Open(); err != nil {
		log.Fatal("Error with database", err)
	}
	srv := &http.Server{
		Addr:    ":8080",
		Handler: app.Routes(),
	}
	srv.ListenAndServe()
}
