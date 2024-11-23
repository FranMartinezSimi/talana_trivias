package questionsrepository

import (
	"context"
	"errors"
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
				// esto le solicite ayuda a claude, tenia respuestas duplicadas y no sabia como solucionarlo
				log.Errorf("all options must have text")
				return errors.New("all options must have text")
			}
		}

		if err := tx.Create(question).Error; err != nil {
			log.Error("Error creating question")
			return err
		}

		if question.CorrectOption > 0 && len(question.Options) > 0 {
			if int(question.CorrectOption) >= len(question.Options) {
				log.Errorf("the correct option index is invalid")
				return errors.New("the correct option index is invalid")
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
	log.Info("Starting full text search with query: ", query)

	var questions []models.Question

	err := q.db.Preload("Options").
		Select("questions.*, "+
			"ts_rank(to_tsvector('english', question), plainto_tsquery(?)) as question_rank, "+
			"COALESCE((SELECT MAX(ts_rank(to_tsvector('english', text), plainto_tsquery(?))) "+
			"FROM options WHERE options.question_id = questions.id), 0) as option_rank",
			query, query).
		Where("to_tsvector('english', question) @@ plainto_tsquery(?) OR EXISTS "+
			"(SELECT 1 FROM options WHERE options.question_id = questions.id AND to_tsvector('english', text) @@ plainto_tsquery(?))",
			query, query).
		Order("question_rank + option_rank DESC").
		Find(&questions).Error

	if err != nil {
		log.WithError(err).Error("Error performing full text search")
		return nil, err
	}

	log.Info("Full text search completed, found ", len(questions), " results")
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
