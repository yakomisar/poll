package main

import (
	"log"
	"poll_service/service"
)

func main() {
	// Создаем структуру для нашего приложения
	// содержит gorm.DB с методами
	app := service.Service{}
	if err := app.Open(); err != nil {
		log.Fatal("Error with database", err)
	}

}
