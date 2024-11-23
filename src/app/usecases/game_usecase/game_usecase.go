package game_usecase

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"talana_prueba_tecnica/src/entity/models"
	"talana_prueba_tecnica/src/entity/requests"
	"talana_prueba_tecnica/src/entity/responses"
	"talana_prueba_tecnica/src/infraestructure/repository/game_repository"
	questionsrepository "talana_prueba_tecnica/src/infraestructure/repository/questions_repository"
	triviarepository "talana_prueba_tecnica/src/infraestructure/repository/trivia_repository"
)

type GameUseCase struct {
	repository   game_repository.GameRepositoryInterface
	questionRepo questionsrepository.QuestionRepositoryInterface
	triviaRepo   triviarepository.TriviaRepositoryInterface
}

func NewGameUseCase(
	repository game_repository.GameRepositoryInterface,
	questionRepo questionsrepository.QuestionRepositoryInterface,
	triviaRepository triviarepository.TriviaRepositoryInterface,
) *GameUseCase {
	return &GameUseCase{
		repository:   repository,
		questionRepo: questionRepo,
		triviaRepo:   triviaRepository,
	}
}

func (u *GameUseCase) GetQuestionsForTrivia(ctx context.Context, triviaID uint) ([]responses.QuestionResponse, error) {
	log := logrus.WithContext(ctx)
	log.Infof("Getting questions for trivia ID %d usecase", triviaID)

	questions, err := u.repository.GetQuestionsForTrivia(ctx, triviaID)
	if err != nil {
		log.WithError(err).Error("Error getting questions for trivia")
		return nil, err
	}

	var response []responses.QuestionResponse
	for _, question := range questions {
		var options []responses.OptionResponse
		for _, option := range question.Options {
			options = append(options, responses.OptionResponse{
				ID:     option.ID,
				Option: option.Text,
			})
		}
		response = append(response, responses.QuestionResponse{
			ID:         question.ID,
			Question:   question.Question,
			Options:    options,
			Difficulty: question.Difficulty,
		})
	}

	log.Info("Questions retrieved successfully")
	return response, nil
}

func (u *GameUseCase) SubmitAnswers(ctx context.Context, triviaID uint, req *requests.SubmitAnswersRequest) (responses.SubmitAnswersResponse, error) {
	log := logrus.WithContext(ctx)
	log.Infof("Submitting answers for trivia ID %d usecase", triviaID)

	if len(req.Responses) == 0 {
		log.Error("No answers provided")
		return responses.SubmitAnswersResponse{}, errors.New("no answers provided")
	}

	var score int
	var correctAnswers int
	for _, response := range req.Responses {
		question, err := u.questionRepo.FindByID(ctx, response.QuestionID)
		if err != nil {
			log.WithError(err).Errorf("Question ID %d not found", response.QuestionID)
			return responses.SubmitAnswersResponse{}, errors.New("invalid question in responses")
		}

		isCorrect := question.CorrectOption == response.SelectedOption
		if isCorrect {
			correctAnswers++
			switch question.Difficulty {
			case "facil":
				score += 1
			case "medio":
				score += 2
			case "dificil":
				score += 3
			}
		}

		playerAnswer := &models.Answer{
			QuestionID:     response.QuestionID,
			SelectedOption: response.SelectedOption,
			IsCorrect:      isCorrect,
		}
		if err := u.repository.SaveAnswer(ctx, playerAnswer); err != nil {
			log.WithError(err).Errorf("Error saving answer for question ID %d", response.QuestionID)
			return responses.SubmitAnswersResponse{}, err
		}
	}

	participation := &models.Participation{
		UserID:   req.UserID,
		TriviaID: triviaID,
		Score:    score,
	}
	if err := u.triviaRepo.SaveParticipation(ctx, participation); err != nil {
		log.WithError(err).Error("Error saving participation")
		return responses.SubmitAnswersResponse{}, err
	}

	log.Infof("Answers submitted successfully with score: %d", score)
	return responses.SubmitAnswersResponse{
		TriviaID:       triviaID,
		UserID:         req.UserID,
		CorrectAnswers: correctAnswers,
		TotalQuestions: len(req.Responses),
		Score:          score,
	}, nil
}
