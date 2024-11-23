package game_usecase

import (
	"context"
	"talana_prueba_tecnica/src/entity/requests"
	"talana_prueba_tecnica/src/entity/responses"
)

type GameUseCaseInterface interface {
	GetQuestionsForTrivia(ctx context.Context, triviaID uint) ([]responses.QuestionResponse, error)
	SubmitAnswers(ctx context.Context, triviaID uint, req *requests.SubmitAnswersRequest) (responses.SubmitAnswersResponse, error)
}
