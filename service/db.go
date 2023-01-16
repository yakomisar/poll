package service

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"poll_service/models"
)

func (s *Service) createPoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := models.Poll{}
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := s.DB.Model(models.Poll{}).Preload("Choice").Create(&p).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Successfully created.")
}

func (s *Service) vote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	d := models.Choice{}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := s.DB.Model(&d).Where("name = ? AND poll_id = ?", d.Name, d.PollID).Update("votes", gorm.Expr("votes + ?", 1)).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(fmt.Sprintf("Successfully voted to %v", d.Name))
}

func (s *Service) getResult(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// подготовим пустую структуру для ответа
	p := models.Poll{}
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := s.DB.Preload("Choice").First(&p).Where("id = ?", p.ID).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(&p)
}
