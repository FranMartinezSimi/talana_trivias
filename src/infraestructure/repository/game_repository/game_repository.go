package game_repository

import (
	"context"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"talana_prueba_tecnica/src/entity/models"
)

type GameRepository struct {
	db *gorm.DB
}

func NewGameRepository(db *gorm.DB) *GameRepository {
	return &GameRepository{db: db}
}

func (r *GameRepository) GetQuestionsForTrivia(ctx context.Context, triviaID uint) ([]models.Question, error) {
	log := logrus.WithContext(ctx)
	log.Infof("GetQuestionForTrivia: triviaID: %d", triviaID)

	var questions []models.Question
	err := r.db.Preload("Options").Joins("JOIN trivia_questions ON trivia_questions.question_id = questions.id").
		Where("trivia_questions.trivia_id = ?", triviaID).Find(&questions).Error
	if err != nil {
		log.Errorf("GetQuestionForTrivia: %v", err)
		return nil, err
	}

	log.Info("GetQuestionForTrivia: success")
	return questions, nil
}

func (r *GameRepository) SaveAnswer(ctx context.Context, answer *models.Answer) error {
	log := logrus.WithContext(ctx)
	log.Infof("SaveAnswer: answer: %v", answer)

	err := r.db.Create(answer).Error
	if err != nil {
		log.Errorf("SaveAnswer: %v", err)
		return err
	}

	log.Info("SaveAnswer: success")
	return nil
}

func (r *GameRepository) GetRankingForTrivia(ctx context.Context, triviaID uint) ([]models.Ranking, error) {
	log := logrus.WithContext(ctx)
	log.Infof("GetRankingForTrivia: triviaID: %d", triviaID)

	var ranking []models.Ranking
	err := r.db.Table("participations").Select("user_id, sum(score) as total_score").
		Where("trivia_id = ?", triviaID).Group("user_id").Order("total_score desc").Find(&ranking).Error

	if err != nil {
		log.Errorf("GetRankingForTrivia: %v", err)
		return nil, err
	}

	log.Info("GetRankingForTrivia: success")
	return ranking, nil
}
