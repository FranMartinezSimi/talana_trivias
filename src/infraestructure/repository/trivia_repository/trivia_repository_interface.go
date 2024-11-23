package triviarepository

import (
	"context"
	"talana_prueba_tecnica/src/entity/models"
)

type TriviaRepositoryInterface interface {
	CreateTrivia(ctx context.Context, trivia *models.Trivia) error
	FindAll(ctx context.Context) ([]models.Trivia, error)
	FindByID(ctx context.Context, id uint) (models.Trivia, error)
	UpdateTrivia(ctx context.Context, trivia *models.Trivia, id uint) error
	DeleteTrivia(ctx context.Context, id uint) error
	FindQuestionByID(ctx context.Context, questionID uint) (models.Question, error)
	SaveParticipation(ctx context.Context, participation *models.Participation) error
	GetUserScore(ctx context.Context, triviaID, userID uint) (models.Participation, error)
	AssignUserToTrivia(ctx context.Context, TriviaID, UserID uint) error
	GetTriviaRanking(ctx context.Context, triviaID uint) ([]models.Ranking, error)
}
