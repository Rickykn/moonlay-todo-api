package repositories

import (
	"github.com/Rickykn/moonlay-todo-api/models"
	"gorm.io/gorm"
)

type TodoRepository interface {
	Create(title, description, file string) (*models.Todo, error)
}

type todoRepository struct {
	db *gorm.DB
}

type TRConfig struct {
	DB *gorm.DB
}

func NewTodoRepository(c *TRConfig) TodoRepository {
	return &todoRepository{
		db: c.DB,
	}
}

func (t *todoRepository) Create(title, description, file string) (*models.Todo, error) {
	newTodo := &models.Todo{
		Title:       title,
		Description: description,
		File:        file,
	}
	result := t.db.Create(&newTodo)

	return newTodo, result.Error
}
