package triviausecase

import (
	"context"
	"talana_prueba_tecnica/src/entity/requests"
	"talana_prueba_tecnica/src/entity/responses"
)

type TriviaUseCaseInterface interface {
	FindAll(ctx context.Context) ([]responses.TriviaResponse, error)
	FindByID(ctx context.Context, id uint) (responses.TriviaResponse, error)
	CreateTrivia(ctx context.Context, req *requests.CreateTriviaRequest) error
	UpdateTrivia(ctx context.Context, req *requests.CreateTriviaRequest, id uint) error
	DeleteTrivia(ctx context.Context, id uint) error
}
