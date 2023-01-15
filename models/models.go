package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

// Создадим структуры для таблиц сервиса голосований
// Структура для конкретного голосования
type Poll struct {
	ID        int        `gorm:"primary_key"`
	Name      string     `json:"name"`
	Choice    []Choice   `gorm:"foreignKey:PollID" json:"choice"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

// Структура для вариантов ответа
type Choice struct {
	PollID    int        `json:"id"`
	Name      string     `json:"name"`
	Votes     int        `json:"votes"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

func (c *Choice) BeforeCreate(tx *gorm.DB) (err error) {
	cols := []clause.Column{}
	for _, field := range tx.Statement.Schema.PrimaryFields {
		cols = append(cols, clause.Column{Name: field.DBName})
	}
	tx.Statement.AddClause(clause.OnConflict{
		Columns:   cols,
		DoNothing: true,
	})
	return nil
}
