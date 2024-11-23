package questionsrepository

import (
	"context"
	"talana_prueba_tecnica/src/entity/models"
)

type QuestionRepositoryInterface interface {
	CreateQuestion(ctx context.Context, question *models.Question) error
	FindAll(ctx context.Context) ([]models.Question, error)
	FindByID(ctx context.Context, id uint) (*models.Question, error)
	FullTextSearch(ctx context.Context, query string) ([]models.Question, error)
	UpdateQuestion(ctx context.Context, question *models.Question, id uint) error
	DeleteQuestion(ctx context.Context, id uint) error
}
