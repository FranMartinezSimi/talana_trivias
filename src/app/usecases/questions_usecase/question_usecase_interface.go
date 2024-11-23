package questionsusecase

import (
	"context"
	"talana_prueba_tecnica/src/entity/requests"
	"talana_prueba_tecnica/src/entity/responses"
)

type QuestionUseCaseInterface interface {
	FindAll(ctx context.Context) ([]responses.QuestionResponse, error)
	FindByID(ctx context.Context, id uint) (responses.QuestionResponse, error)
	CreateQuestion(ctx context.Context, req *requests.CreateQuestionRequest) error
	UpdateQuestion(ctx context.Context, req *requests.CreateQuestionRequest, id uint) error
	DeleteQuestion(ctx context.Context, id uint) error
	FullTextSearch(ctx context.Context, search string) ([]responses.QuestionResponse, error)
}
