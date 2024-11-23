package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"strconv"
	gameusecase "talana_prueba_tecnica/src/app/usecases/game_usecase"
	"talana_prueba_tecnica/src/entity/requests"
)

type GameHandler struct {
	useCase gameusecase.GameUseCaseInterface
}

func NewGameHandler(useCase gameusecase.GameUseCaseInterface) *GameHandler {
	return &GameHandler{
		useCase: useCase,
	}
}

// @Summary Get questions for a trivia
// @Description Retrieve all questions for a specific trivia
// @Tags Games
// @Param id path uint true "Trivia ID"
// @Produce json
// @Success 200 {object} []responses.QuestionResponse "Questions for the trivia"
// @Failure 400 {object} map[string]interface{} "Invalid trivia ID"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /games/trivias/{id}/questions [get]
func (h *GameHandler) GetQuestionsForTrivia(ctx *fiber.Ctx) error {
	log := logrus.WithContext(ctx.Context())
	log.Info("Get questions for trivia handler")

	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		log.Errorf("Invalid trivia ID: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid trivia ID"})
	}

	questions, err := h.useCase.GetQuestionsForTrivia(ctx.Context(), uint(id))
	if err != nil {
		log.Errorf("Error getting questions for trivia: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	log.Info("Questions for trivia retrieved successfully")
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": questions})
}

// @Summary Submit answers for a trivia
// @Description Submit answers for a specific trivia and calculate the user's score
// @Tags Games
// @Accept json
// @Produce json
// @Param id path uint true "Trivia ID"
// @Param answers body requests.SubmitAnswersRequest true "User answers"
// @Success 200 {object} responses.SubmitAnswersResponse "User score and details"
// @Failure 400 {object} map[string]interface{} "Invalid request or trivia ID"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /games/trivias/{id}/answers [post]
func (h *GameHandler) SubmitAnswers(ctx *fiber.Ctx) error {
	log := logrus.WithContext(ctx.Context())
	log.Info("Submit answers handler")

	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		log.Errorf("Invalid trivia ID: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid trivia ID"})
	}

	var req requests.SubmitAnswersRequest
	if err := ctx.BodyParser(&req); err != nil {
		log.Errorf("Error parsing request: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	response, err := h.useCase.SubmitAnswers(ctx.Context(), uint(id), &req)
	if err != nil {
		log.Errorf("Error submitting answers: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	log.Info("Answers submitted successfully")
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": response})
}
