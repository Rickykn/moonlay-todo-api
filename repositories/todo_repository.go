package repositories

import (
	"github.com/Rickykn/moonlay-todo-api/helpers"
	"github.com/Rickykn/moonlay-todo-api/models"
	"gorm.io/gorm"
)

type TodoRepository interface {
	Create(title, description, file string) (*models.Todo, error)
	FindAllTodo(query *helpers.Query) ([]*models.Todo, error)
	FindOneTodoByid(id int) (*models.Todo, int, error)
	FindAllTodoWithSubTodo() ([]*models.Todo, error)
	Delete(id int) (int, error)
	Update(newTodo *models.Todo) (*models.Todo, int, error)
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

func (t *todoRepository) FindAllTodo(query *helpers.Query) ([]*models.Todo, error) {
	var todos []*models.Todo

	result := t.db.
		Offset((query.Page - 1) * query.Limit).
		Limit(query.Limit).
		Find(&todos)

	return todos, result.Error
}

func (t *todoRepository) FindOneTodoByid(id int) (*models.Todo, int, error) {
	var todo *models.Todo
	result := t.db.Where("id = ?", id).First(&todo)

	return todo, int(result.RowsAffected), result.Error
}

func (t *todoRepository) FindAllTodoWithSubTodo() ([]*models.Todo, error) {
	var todos []*models.Todo

	result := t.db.Preload("Sub_todo").Find(&todos)

	return todos, result.Error
}

func (t *todoRepository) Delete(id int) (int, error) {
	var todo *models.Todo

	result := t.db.Where("id = ?", id).Delete(&todo)
	return int(result.RowsAffected), result.Error
}

func (t *todoRepository) Update(newTodo *models.Todo) (*models.Todo, int, error) {
	result := t.db.Save(newTodo)

	return newTodo, int(result.RowsAffected), result.Error
}
