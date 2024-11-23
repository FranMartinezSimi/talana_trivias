package triviarepository

import (
	"context"
	"talana_prueba_tecnica/src/entity/models"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TriviaRepository struct {
	db *gorm.DB
}

func NewTriviaRepository(db *gorm.DB) *TriviaRepository {
	return &TriviaRepository{db: db}
}

func (r *TriviaRepository) CreateTrivia(ctx context.Context, trivia *models.Trivia) error {
	log := logrus.WithContext(ctx)
	log.Info("Creating trivia")

	err := r.db.Create(trivia)
	if err.Error != nil {
		log.WithError(err.Error).Error("Error creating trivia")
		return err.Error
	}

	log.Info("Trivia created")
	return nil
}

func (r *TriviaRepository) FindAll(ctx context.Context) ([]models.Trivia, error) {
	log := logrus.WithContext(ctx)
	log.Info("Finding all trivias")

	var trivias []models.Trivia
	err := r.db.Preload("Questions").Preload("Users").Find(&trivias)
	if err.Error != nil {
		log.WithError(err.Error).Error("Error finding all trivias")
		return nil, err.Error
	}

	log.Info("Trivias found")
	return trivias, nil
}

func (r *TriviaRepository) FindByID(ctx context.Context, id uint) (models.Trivia, error) {
	log := logrus.WithContext(ctx)
	log.Infof("Finding trivia by ID: %d", id)

	var trivia models.Trivia
	err := r.db.Preload("Questions").Preload("Users").First(&trivia, id)
	if err.Error != nil {
		log.WithError(err.Error).Error("Error finding trivia by ID")
		return models.Trivia{}, err.Error
	}

	log.Info("Trivia found")
	return trivia, nil
}

func (r *TriviaRepository) UpdateTrivia(ctx context.Context, trivia *models.Trivia, id uint) error {
	log := logrus.WithContext(ctx)
	log.Infof("Updating trivia with ID: %d", id)

	err := r.db.Model(&models.Trivia{}).Where("id = ?", id).Updates(trivia)
	if err.Error != nil {
		log.WithError(err.Error).Error("Error updating trivia")
		return err.Error
	}

	log.Info("Trivia updated")
	return nil
}

func (r *TriviaRepository) DeleteTrivia(ctx context.Context, id uint) error {
	log := logrus.WithContext(ctx)
	log.Infof("Deleting trivia with ID: %d", id)

	err := r.db.Delete(&models.Trivia{}, id)
	if err.Error != nil {
		log.WithError(err.Error).Error("Error deleting trivia")
		return err.Error
	}

	log.Info("Trivia deleted")
	return nil
}

func (r *TriviaRepository) SaveParticipation(ctx context.Context, participation *models.Participation) error {
	log := logrus.WithContext(ctx)
	log.Info("Saving participation")

	err := r.db.Create(participation)
	if err.Error != nil {
		log.WithError(err.Error).Error("Error saving participation")
		return err.Error
	}

	log.Info("Participation saved")
	return nil
}

func (r *TriviaRepository) GetUserScore(ctx context.Context, triviaID, userID uint) (models.Participation, error) {
	log := logrus.WithContext(ctx)
	log.Infof("Getting user score for trivia ID: %d and user ID: %d", triviaID, userID)

	var participation models.Participation
	err := r.db.Preload("Answers").Where("trivia_id = ? AND user_id = ?", triviaID, userID).First(&participation).Error
	if err != nil {
		log.WithError(err).Error("Error getting user score")
		return models.Participation{}, err
	}

	log.Info("User score retrieved successfully")
	return participation, nil
}

func (r *TriviaRepository) FindQuestionByID(ctx context.Context, questionID uint) (models.Question, error) {
	log := logrus.WithContext(ctx)
	log.Infof("Finding question by ID: %d", questionID)

	var question models.Question
	err := r.db.Preload("Options").First(&question, questionID).Error
	if err != nil {
		log.WithError(err).Error("Error finding question by ID")
		return models.Question{}, err
	}

	log.Info("Question found successfully")
	return question, nil
}

func (r *TriviaRepository) AssignUserToTrivia(ctx context.Context, triviaID uint, userID uint) error {
	log := logrus.WithContext(ctx)
	log.Infof("Assigning user ID %d to trivia ID %d", userID, triviaID)

	triviaUser := models.TriviaUser{
		TriviaID: triviaID,
		UserID:   userID,
	}

	err := r.db.Create(&triviaUser).Error
	if err != nil {
		log.WithError(err).Error("Error assigning user to trivia")
		return err
	}

	log.Info("User assigned to trivia successfully")
	return nil
}

func (r *TriviaRepository) GetTriviaRanking(ctx context.Context, triviaID uint) ([]models.Ranking, error) {
	log := logrus.WithContext(ctx)
	log.Infof("Getting ranking for trivia ID %d", triviaID)

	var rankings []models.Ranking
	err := r.db.Table("participations").
		Select("user_id, SUM(score) as total_score").
		Where("trivia_id = ?", triviaID).
		Group("user_id").
		Order("total_score DESC").
		Find(&rankings).Error

	if err != nil {
		log.WithError(err).Error("Error retrieving trivia ranking")
		return nil, err
	}

	log.Info("Trivia ranking retrieved successfully")
	return rankings, nil
}
