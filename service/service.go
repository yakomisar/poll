package service

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

type Service struct {
	DB *gorm.DB
}

func (s *Service) Open() error {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error when loading .env file")
		return err
	}
	// Пробуем открыть БД, используя файл окружения
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE")), &gorm.Config{})
	if err != nil {
		log.Fatal("Error while openning database")
		return err
	}
	// Записываем БД в нашу структуру
	s.DB = db
	// Запускаем наш роутер, который отслеживает работу с запросами
	s.Routes()
	// Drop table if exists (will ignore or delete foreign key constraints when dropping)
	s.DB.Migrator().DropTable(&models.Poll{}, &models.Choice{})
	return s.DB.AutoMigrate(&models.Poll{}, &models.Choice{})
}

func (s *Service) Routes() http.Handler {
	r := mux.NewRouter()
	// Создадим маршруты
	// 1. Создание голосования
	r.HandleFunc("/api/createpoll", s.createPoll).Methods("POST")
	// 2. Голосование
	r.HandleFunc("/api/poll", s.vote).Methods("POST")
	// 3. Получение результатов
	r.HandleFunc("/api/getresult", s.getResult).Methods("POST")
	return r
}
