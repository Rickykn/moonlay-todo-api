package dtos

type ResponeTodoDTO struct {
	ID          int    `json:"id"`
	Title       string `json:"tittle"`
	Description string `json:"description"`
	File        string `json:"file"`
}
