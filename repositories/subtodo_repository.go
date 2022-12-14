package repositories

import (
	"github.com/Rickykn/moonlay-todo-api/models"
	"gorm.io/gorm"
)

type SubtodoRepository interface {
	Create(title, description, file string, id int) (*models.SubTodo, error)
	FindAllByTodoId(id int) ([]*models.SubTodo, error)
	Delete(id int) (int, error)
	Update(newSubtodo *models.SubTodo) (*models.SubTodo, int, error)
	FindSubtodoById(id int) (*models.SubTodo, int, error)
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

func (s *subtodoRepository) FindAllByTodoId(id int) ([]*models.SubTodo, error) {
	var subtodos []*models.SubTodo

	result := s.db.Where("todo_id=?", id).Find(&subtodos)

	return subtodos, result.Error
}

func (s *subtodoRepository) Delete(id int) (int, error) {
	var subtodo *models.SubTodo

	result := s.db.Where("id = ?", id).Delete(&subtodo)
	return int(result.RowsAffected), result.Error
}

func (s *subtodoRepository) Update(newSubtodo *models.SubTodo) (*models.SubTodo, int, error) {
	result := s.db.Save(newSubtodo)

	return newSubtodo, int(result.RowsAffected), result.Error
}

func (s *subtodoRepository) FindSubtodoById(id int) (*models.SubTodo, int, error) {
	var subtodo *models.SubTodo
	result := s.db.Where("id = ?", id).First(&subtodo)

	return subtodo, int(result.RowsAffected), result.Error
}
