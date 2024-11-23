package game_repository

import (
	"context"
	"talana_prueba_tecnica/src/entity/models"
)

type GameRepositoryInterface interface {
	GetQuestionsForTrivia(ctx context.Context, triviaID uint) ([]models.Question, error)
	SaveAnswer(ctx context.Context, answer *models.Answer) error
	GetRankingForTrivia(ctx context.Context, triviaID uint) ([]models.Ranking, error)
}
