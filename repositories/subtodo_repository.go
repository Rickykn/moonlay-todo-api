package repositories

import (
	"github.com/Rickykn/moonlay-todo-api/models"
	"gorm.io/gorm"
)

type SubtodoRepository interface {
	Create(title, description, file string, id int) (*models.SubTodo, error)
}

type subtodoRepository struct {
	db *gorm.DB
}

type SRConfig struct {
	DB *gorm.DB
}

func NewSubtodoRepository(c *SRConfig) SubtodoRepository {
	return &subtodoRepository{
		db: c.DB,
	}
}

func (s *subtodoRepository) Create(title, description, file string, id int) (*models.SubTodo, error) {
	newSubTodo := &models.SubTodo{
		Title:       title,
		Description: description,
		File:        file,
		Todo_id:     &id,
	}
	result := s.db.Joins("Todo").Where("todo_id = ?", id).Create(&newSubTodo)

	return newSubTodo, result.Error
}
