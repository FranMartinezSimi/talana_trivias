package questionsusecase

import (
	"context"
	"errors"
	"talana_prueba_tecnica/src/entity/models"
	"talana_prueba_tecnica/src/entity/requests"
	"talana_prueba_tecnica/src/entity/responses"
	questionsrepository "talana_prueba_tecnica/src/infraestructure/repository/questions_repository"

	"github.com/sirupsen/logrus"
)

type QuestionsUseCase struct {
	repository questionsrepository.QuestionRepositoryInterface
}

func NewQuestionsUseCase(repository questionsrepository.QuestionRepositoryInterface) *QuestionsUseCase {
	return &QuestionsUseCase{
		repository: repository,
	}

}
func (u *QuestionsUseCase) FindAll(ctx context.Context) ([]responses.QuestionResponse, error) {
	log := logrus.WithContext(ctx)
	log.Info("Get all questions usecase")

	result, err := u.repository.FindAll(ctx)
	if err != nil {
		log.Errorf("Error: %v", err)
		return nil, err
	}

	log.Info("Questions found")
	var questionsList []responses.QuestionResponse

	for _, question := range result {
		var optionsList []responses.OptionResponse
		for _, option := range question.Options {
			optionsList = append(optionsList, responses.OptionResponse{
				ID:     option.ID,
				Option: option.Text,
			})
		}

		responseQuestion := responses.QuestionResponse{
			ID:            question.ID,
			Question:      question.Question,
			CorrectOption: question.CorrectOption,
			Options:       optionsList,
			Difficulty:    question.Difficulty,
		}

		questionsList = append(questionsList, responseQuestion)
	}

	return questionsList, nil
}

func (u *QuestionsUseCase) FindByID(ctx context.Context, id uint) (responses.QuestionResponse, error) {
	log := logrus.WithContext(ctx)
	log.Infof("Get question by ID: %d usecase", id)

	result, err := u.repository.FindByID(ctx, id)
	if err != nil {
		log.Errorf("Error: %v", err)
		return responses.QuestionResponse{}, err
	}

	log.Info("Question found")

	var optionsList []responses.OptionResponse
	for _, option := range result.Options {
		optionsList = append(optionsList, responses.OptionResponse{
			ID:     option.ID,
			Option: option.Text,
		})
	}

	responseQuestion := responses.QuestionResponse{
		ID:            result.ID,
		Question:      result.Question,
		CorrectOption: result.CorrectOption,
		Options:       optionsList,
		Difficulty:    result.Difficulty,
	}

	return responseQuestion, nil
}

func (u *QuestionsUseCase) CreateQuestion(ctx context.Context, req *requests.CreateQuestionRequest) error {
	log := logrus.WithContext(ctx)
	log.Info("Creating question in usecase")

	if len(req.Options) < 2 {
		log.Errorf("at least two options are required")
		return errors.New("at least two options are required")
	}

	if req.CorrectOption >= len(req.Options) || req.CorrectOption < 0 {
		log.Errorf("invalid correct option index")
		return errors.New("invalid correct option index")
	}

	var options []models.Option

	for _, opt := range req.Options {
		options = append(options, models.Option{Text: opt})
	}

	question := &models.Question{
		Question:      req.Question,
		Difficulty:    req.Difficulty,
		Points:        req.Points,
		CorrectOption: uint(req.CorrectOption),
		Options:       options,
	}

	err := u.repository.CreateQuestion(ctx, question)
	if err != nil {
		log.WithError(err).Error("Error creating question in repository")
		return err
	}

	log.Info("Question created successfully")
	return nil
}

func (u *QuestionsUseCase) UpdateQuestion(ctx context.Context, req *requests.CreateQuestionRequest, id uint) error {
	log := logrus.WithContext(ctx)
	log.Info("Updating question in usecase")

	if len(req.Options) < 2 {
		log.Errorf("at least two options are required")
		return errors.New("at least two options are required")
	}

	if req.CorrectOption >= len(req.Options) || req.CorrectOption < 0 {
		log.Errorf("invalid correct option index")
		return errors.New("invalid correct option index")
	}

	question := &models.Question{
		Question:      req.Question,
		Difficulty:    req.Difficulty,
		Points:        req.Points,
		CorrectOption: uint(req.CorrectOption),
		Options:       make([]models.Option, len(req.Options)),
	}

	for _, opt := range req.Options {
		question.Options = append(question.Options, models.Option{Text: opt})
	}

	err := u.repository.UpdateQuestion(ctx, question, id)
	if err != nil {
		log.WithError(err).Error("Error updating question in repository")
		return err
	}

	log.Info("Question updated successfully")
	return nil
}

func (u *QuestionsUseCase) DeleteQuestion(ctx context.Context, id uint) error {
	log := logrus.WithContext(ctx)
	log.Info("Delete question usecase")

	questionExists, err := u.repository.FindByID(ctx, id)
	if err != nil {
		log.Errorf("Error: %v", err)
		return err
	}

	if questionExists.ID == 0 {
		log.Errorf("Question not found")
		return errors.New("Question not found")
	}

	err = u.repository.DeleteQuestion(ctx, id)
	if err != nil {
		log.Errorf("Error: %v", err)
		return err
	}

	log.Info("Question deleted successfully")
	return nil
}

func (u *QuestionsUseCase) FullTextSearch(ctx context.Context, query string) ([]responses.QuestionResponse, error) {
	log := logrus.WithContext(ctx)
	log.Infof("Performing full text search with query: %s", query)

	result, err := u.repository.FullTextSearch(ctx, query)
	if err != nil {
		log.WithError(err).Error("Error during full text search in repository")
		return nil, err
	}

	log.Infof("Questions found for query: %s", query)
	var questionsList []responses.QuestionResponse

	for _, question := range result {
		var optionsList []responses.OptionResponse
		for _, option := range question.Options {
			optionsList = append(optionsList, responses.OptionResponse{
				ID:     option.ID,
				Option: option.Text,
			})
		}

		responseQuestion := responses.QuestionResponse{
			ID:            question.ID,
			Question:      question.Question,
			CorrectOption: question.CorrectOption,
			Options:       optionsList,
			Difficulty:    question.Difficulty,
		}

		questionsList = append(questionsList, responseQuestion)
	}

	return questionsList, nil
}
