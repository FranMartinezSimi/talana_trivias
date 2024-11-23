package questionsrepository

import (
	"context"
	"fmt"
	"talana_prueba_tecnica/src/entity/models"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type QuestionRepository struct {
	db *gorm.DB
}

func NewQuestionRepository(gorm *gorm.DB) *QuestionRepository {
	return &QuestionRepository{
		db: gorm,
	}
}

func (q *QuestionRepository) CreateQuestion(ctx context.Context, question *models.Question) error {
	log := logrus.WithContext(ctx)
	log.Info("creating question in repository")

	return q.db.Transaction(func(tx *gorm.DB) error {
		for _, option := range question.Options {
			if option.Text == "" {
				return fmt.Errorf("todas las opciones deben tener texto")
			}
		}

		if err := tx.Create(question).Error; err != nil {
			log.Error("Error creating question")
			return err
		}

		if question.CorrectOption > 0 && len(question.Options) > 0 {
			if int(question.CorrectOption) >= len(question.Options) {
				return fmt.Errorf("índice de opción correcta inválido")
			}

			question.CorrectOption = question.Options[question.CorrectOption-1].ID
			if err := tx.Save(question).Error; err != nil {
				log.Error("Error updating question")
				return err
			}
		}
		return nil
	})
}

func (q *QuestionRepository) FindAll(ctx context.Context) ([]models.Question, error) {
	log := logrus.WithContext(ctx)
	log.Println("finding all questions")

	var questions []models.Question

	log.Info("finding all questions")

	res := q.db.WithContext(ctx).Preload("Options").Find(&questions)
	if res.Error != nil {
		log.Error("Error finding all questions")
		return nil, res.Error
	}

	log.WithError(res.Error).Info("questions found")
	return questions, nil
}

func (q *QuestionRepository) FindByID(ctx context.Context, id uint) (*models.Question, error) {
	log := logrus.WithContext(ctx)
	log.Println("finding question by id")

	var question models.Question

	log.Info("finding question by id")

	res := q.db.WithContext(ctx).Preload("Options").First(&question, id)
	if res.Error != nil {
		log.Error("Error finding question by id")
		return nil, res.Error
	}

	log.WithError(res.Error).Info("question found")
	return &question, nil
}

func (q *QuestionRepository) FullTextSearch(ctx context.Context, query string) ([]models.Question, error) {
	log := logrus.WithContext(ctx)
	log.Println("full text search")

	var questions []models.Question

	err := q.db.Preload("Options").
		Where("to_tsvector('english', text) @@ plainto_tsquery(?) OR EXISTS "+
			"(SELECT 1 FROM options WHERE options.question_id = questions.id AND to_tsvector('english', text) @@ plainto_tsquery(?))", query, query).
		Find(&questions).Error

	if err != nil {
		return nil, err
	}

	return questions, nil
}

func (q *QuestionRepository) UpdateQuestion(ctx context.Context, question *models.Question, id uint) error {
	log := logrus.WithContext(ctx)
	log.Info("Updating question")

	err := q.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.Question{}).Where("id = ?", id).Updates(question).Error; err != nil {
			log.WithError(err).Error("Error updating question")
			return err
		}

		log.Info("Question updated successfully")
		return nil
	})

	if err != nil {
		log.WithError(err).Error("Transaction failed while updating question")
		return err
	}

	log.Info("Transaction completed successfully")
	return nil
}

func (q *QuestionRepository) DeleteQuestion(ctx context.Context, id uint) error {
	log := logrus.WithContext(ctx)
	log.Info("deleting question")

	err := q.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&models.Question{}, id).Error; err != nil {
			log.Error("Error deleting question")
			return err
		}

		return nil
	})

	if err != nil {
		log.WithError(err).Error("Error deleting question")
		return err
	}

	log.Info("question deleted")
	return nil
}
