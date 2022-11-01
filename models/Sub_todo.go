package models

import (
	"time"

	"gorm.io/gorm"
)

type SubTodo struct {
	ID          int    `json:"id" gorm:"primary_key"`
	Title       string `json:"tittle" gorm:"size:100;not null"`
	Description string `json:"description" gorm:"size:1000;not null"`
	File        string `json:"file"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
	Todo_id     *int `json:"todo_id"`
	Todo        Todo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
