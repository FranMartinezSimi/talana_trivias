package triviausecase

import (
	"context"
	"errors"
	"talana_prueba_tecnica/src/entity/models"
	"talana_prueba_tecnica/src/entity/requests"
	"talana_prueba_tecnica/src/entity/responses"
	triviarepository "talana_prueba_tecnica/src/infraestructure/repository/trivia_repository"

	"github.com/sirupsen/logrus"
)

type TriviaUseCase struct {
	triviaRepository triviarepository.TriviaRepositoryInterface
}

func NewTriviaUseCase(triviaRepository triviarepository.TriviaRepositoryInterface) *TriviaUseCase {
	return &TriviaUseCase{
		triviaRepository: triviaRepository,
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

func (u *TriviaUseCase) SaveParticipation(ctx context.Context, req *requests.SaveParticipationRequest) error {
	log := logrus.WithContext(ctx)
	log.Info("Saving participation usecase")

	var score int
	var answers []models.Answer

	for _, answer := range req.Answers {
		question, err := u.triviaRepository.FindQuestionByID(ctx, answer.QuestionID)
		if err != nil {
			log.WithError(err).Errorf("Error finding question ID: %d", answer.QuestionID)
			return err
		}

		isCorrect := question.CorrectOption == answer.SelectedOption
		if isCorrect {
			switch question.Difficulty {
			case "facil":
				score += 1
			case "medio":
				score += 2
			case "dificil":
				score += 3
			}
		}

		answers = append(answers, models.Answer{
			QuestionID:     answer.QuestionID,
			SelectedOption: answer.SelectedOption,
			IsCorrect:      isCorrect,
		})
	}

	participation := &models.Participation{
		UserID:   req.UserID,
		TriviaID: req.TriviaID,
		Score:    score,
		Answers:  answers,
	}

	err := u.triviaRepository.SaveParticipation(ctx, participation)
	if err != nil {
		log.WithError(err).Error("Error saving participation in repository")
		return err
	}

	log.Info("Participation saved successfully")
	return nil
}

func (u *TriviaUseCase) GetUserScore(ctx context.Context, triviaID, userID uint) (responses.UserScoreResponse, error) {
	log := logrus.WithContext(ctx)
	log.Infof("Getting user score for trivia ID: %d and user ID: %d usecase", triviaID, userID)

	participation, err := u.triviaRepository.GetUserScore(ctx, triviaID, userID)
	if err != nil {
		log.WithError(err).Error("Error getting user score in repository")
		return responses.UserScoreResponse{}, err
	}

	correctAnswers := 0
	for _, answer := range participation.Answers {
		if answer.IsCorrect {
			correctAnswers++
		}
	}

	// Crear la respuesta
	response := responses.UserScoreResponse{
		TriviaID:       participation.TriviaID,
		UserID:         participation.UserID,
		Score:          participation.Score,
		CorrectAnswers: correctAnswers,
		TotalQuestions: len(participation.Answers),
	}

	log.Info("User score retrieved successfully")
	return response, nil
}

/*
*** ADDED CODE ***
TODO: Implement the PlayTrivia and SubmitAnswers methods
// */
// func (u *TriviaUseCase) PlayTrivia(ctx context.Context, triviaID uint) (responses.PlayTriviaResponse, error) {
// 	log := logrus.WithContext(ctx)
// 	log.Infof("Preparing trivia for play with ID: %d usecase", triviaID)

// 	trivia, err := u.triviaRepository.FindByID(ctx, triviaID)
// 	if err != nil {
// 		log.WithError(err).Error("Error finding trivia for play in repository")
// 		return responses.PlayTriviaResponse{}, err
// 	}

// 	var questions []responses.QuestionResponse
// 	for _, question := range trivia.Questions {
// 		var options []responses.OptionResponse
// 		for _, option := range question.Options {
// 			options = append(options, responses.OptionResponse{
// 				ID:     option.ID,
// 				Option: option.Text,
// 			})
// 		}
// 		questions = append(questions, responses.QuestionResponse{
// 			ID:         question.ID,
// 			Question:   question.Text,
// 			Options:    options,
// 			Difficulty: question.Difficulty,
// 		})
// 	}

// 	response := responses.PlayTriviaResponse{
// 		ID:          trivia.ID,
// 		Name:        trivia.Name,
// 		Description: trivia.Description,
// 		Questions:   questions,
// 	}

// 	log.Info("Trivia prepared successfully for play")
// 	return response, nil
// }

// func (u *TriviaUseCase) SubmitAnswers(ctx context.Context, triviaID uint, req *requests.SubmitAnswersRequest) (responses.SubmitAnswersResponse, error) {
// 	log := logrus.WithContext(ctx)
// 	log.Infof("Submitting answers for trivia ID: %d usecase", triviaID)

// 	var score int
// 	var correctAnswers int

// 	for _, answer := range req.Responses {
// 		question, err := u.triviaRepository.FindQuestionByID(ctx, answer.QuestionID)
// 		if err != nil {
// 			log.WithError(err).Errorf("Error finding question ID: %d", answer.QuestionID)
// 			return responses.SubmitAnswersResponse{}, err
// 		}

// 		if question.CorrectOption == answer.SelectedOption {
// 			correctAnswers++
// 			switch question.Difficulty {
// 			case "facil":
// 				score += 1
// 			case "medio":
// 				score += 2
// 			case "dificil":
// 				score += 3
// 			}
// 		}
// 	}

// 	participation := &models.Participation{
// 		UserID:   req.UserID,
// 		TriviaID: triviaID,
// 		Score:    score,
// 		Answers:  req.Responses,
// 	}

// 	err := u.triviaRepository.SaveParticipation(ctx, participation)
// 	if err != nil {
// 		log.WithError(err).Error("Error saving participation in repository")
// 		return responses.SubmitAnswersResponse{}, err
// 	}

// 	response := responses.SubmitAnswersResponse{
// 		TriviaID:       triviaID,
// 		UserID:         req.UserID,
// 		CorrectAnswers: correctAnswers,
// 		TotalQuestions: len(req.Responses),
// 		Score:          score,
// 	}

// 	log.Info("Answers submitted successfully")
// 	return response, nil
// }
