package triviausecase

import (
	"context"
	"errors"
	"talana_prueba_tecnica/src/entity/models"
	"talana_prueba_tecnica/src/entity/requests"
	"talana_prueba_tecnica/src/entity/responses"
	triviarepository "talana_prueba_tecnica/src/infraestructure/repository/trivia_repository"
	repository "talana_prueba_tecnica/src/infraestructure/repository/user_repository"

	"github.com/sirupsen/logrus"
)

type TriviaUseCase struct {
	triviaRepository triviarepository.TriviaRepositoryInterface
	userRepository   repository.UserRepositoryInterface
}

func NewTriviaUseCase(
	triviaRepository triviarepository.TriviaRepositoryInterface,
	userRepository repository.UserRepositoryInterface,
) *TriviaUseCase {
	return &TriviaUseCase{
		triviaRepository: triviaRepository,
		userRepository:   userRepository,
	}
}

func (u *TriviaUseCase) CreateTrivia(ctx context.Context, req *requests.CreateTriviaRequest) error {
	log := logrus.WithContext(ctx)
	log.Info("Creating trivia usecase")

	trivia := &models.Trivia{
		Name:        req.Name,
		Description: req.Description,
	}

	for _, questionID := range req.QuestionIDs {
		trivia.Questions = append(trivia.Questions, models.Question{ID: questionID})
	}

	for _, userID := range req.UserIDs {
		trivia.Users = append(trivia.Users, models.UserModel{ID: userID})
	}

	err := u.triviaRepository.CreateTrivia(ctx, trivia)
	if err != nil {
		log.WithError(err).Error("Error creating trivia in repository")
		return err
	}

	log.Info("Trivia created successfully")
	return nil
}

func (u *TriviaUseCase) FindAll(ctx context.Context) ([]responses.TriviaResponse, error) {
	log := logrus.WithContext(ctx)
	log.Info("Finding all trivias usecase")

	trivias, err := u.triviaRepository.FindAll(ctx)
	if err != nil {
		log.WithError(err).Error("Error finding all trivias in repository")
		return nil, err
	}

	var triviaResponses []responses.TriviaResponse
	for _, trivia := range trivias {
		var questionResponses []responses.QuestionResponse
		for _, question := range trivia.Questions {
			var optionResponses []responses.OptionResponse
			for _, option := range question.Options {
				optionResponses = append(optionResponses, responses.OptionResponse{
					ID:     option.ID,
					Option: option.Text,
				})
			}

			questionResponses = append(questionResponses, responses.QuestionResponse{
				ID:         question.ID,
				Question:   question.Question,
				Options:    optionResponses,
				Difficulty: question.Difficulty,
			})
		}

		var userResponses []responses.UserResponse
		for _, user := range trivia.Users {
			userResponses = append(userResponses, responses.UserResponse{
				ID:    user.ID,
				Name:  user.Name,
				Email: user.Email,
			})
		}

		triviaResponses = append(triviaResponses, responses.TriviaResponse{
			ID:          trivia.ID,
			Name:        trivia.Name,
			Description: trivia.Description,
			Questions:   questionResponses,
			Users:       userResponses,
		})
	}

	log.Info("All trivias found successfully")
	return triviaResponses, nil
}

func (u *TriviaUseCase) FindByID(ctx context.Context, id uint) (responses.TriviaResponse, error) {
	log := logrus.WithContext(ctx)
	log.Infof("Finding trivia by ID: %d usecase", id)

	trivia, err := u.triviaRepository.FindByID(ctx, id)
	if err != nil {
		log.WithError(err).Error("Error finding trivia by ID in repository")
		return responses.TriviaResponse{}, err
	}

	var questionResponses []responses.QuestionResponse
	for _, question := range trivia.Questions {
		var optionResponses []responses.OptionResponse
		for _, option := range question.Options {
			optionResponses = append(optionResponses, responses.OptionResponse{
				ID:     option.ID,
				Option: option.Text,
			})
		}

		questionResponses = append(questionResponses, responses.QuestionResponse{
			ID:            question.ID,
			Question:      question.Question,
			CorrectOption: question.CorrectOption,
			Options:       optionResponses,
			Difficulty:    question.Difficulty,
		})
	}

	var userResponses []responses.UserResponse
	for _, user := range trivia.Users {
		userResponses = append(userResponses, responses.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}

	response := responses.TriviaResponse{
		ID:          trivia.ID,
		Name:        trivia.Name,
		Description: trivia.Description,
		Questions:   questionResponses,
		Users:       userResponses,
	}

	log.Info("Trivia found successfully")
	return response, nil
}

func (u *TriviaUseCase) UpdateTrivia(ctx context.Context, req *requests.CreateTriviaRequest, id uint) error {
	log := logrus.WithContext(ctx)
	log.Infof("Updating trivia with ID: %d usecase", id)

	trivia := &models.Trivia{
		Name:        req.Name,
		Description: req.Description,
	}

	for _, questionID := range req.QuestionIDs {
		trivia.Questions = append(trivia.Questions, models.Question{ID: questionID})
	}

	for _, userID := range req.UserIDs {
		trivia.Users = append(trivia.Users, models.UserModel{ID: userID})
	}

	err := u.triviaRepository.UpdateTrivia(ctx, trivia, id)
	if err != nil {
		log.WithError(err).Error("Error updating trivia in repository")
		return err
	}

	log.Info("Trivia updated successfully")
	return nil
}

func (u *TriviaUseCase) DeleteTrivia(ctx context.Context, id uint) error {
	log := logrus.WithContext(ctx)
	log.Infof("Deleting trivia with ID: %d usecase", id)

	_, err := u.triviaRepository.FindByID(ctx, id)
	if err != nil {
		log.WithError(err).Error("Error finding trivia for deletion in repository")
		return errors.New("trivia not found")
	}

	err = u.triviaRepository.DeleteTrivia(ctx, id)
	if err != nil {
		log.WithError(err).Error("Error deleting trivia in repository")
		return err
	}

	log.Info("Trivia deleted successfully")
	return nil
}

func (u *TriviaUseCase) AssignUserToTrivia(ctx context.Context, TriviaID, UserID uint) error {
	log := logrus.WithContext(ctx)
	log.Infof("Assigning user ID: %d to trivia ID: %d usecase", UserID, TriviaID)

	trivia, err := u.triviaRepository.FindByID(ctx, TriviaID)
	if err != nil {
		log.WithError(err).Error("Error finding trivia for assigning user in repository")
		return err
	}

	user, err := u.userRepository.FindByID(ctx, UserID)
	if err != nil {
		log.WithError(err).Error("Error finding user for assigning to trivia in repository")
		return err
	}

	err = u.triviaRepository.AssignUserToTrivia(ctx, trivia.ID, user.ID)
	if err != nil {
		log.WithError(err).Error("Error assigning user to trivia in repository")
		return err
	}

	log.Info("User assigned to trivia successfully")
	return nil
}
