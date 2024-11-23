package handlers

import (
	"strconv"
	questionsusecase "talana_prueba_tecnica/src/app/usecases/questions_usecase"
	"talana_prueba_tecnica/src/entity/requests"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type QuestionHandler struct {
	useCase questionsusecase.QuestionUseCaseInterface
}

func NewQuestionHandler(useCase questionsusecase.QuestionUseCaseInterface) *QuestionHandler {
	return &QuestionHandler{
		useCase: useCase,
	}
}

// @Summary Get all questions
// @Description Retrieve a list of all available questions
// @Tags Questions
// @Produce json
// @Success 200 {object} []responses.QuestionResponse "List of questions"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /questions [get]

func (h *QuestionHandler) GetAllQuestions(ctx *fiber.Ctx) error {
	log := logrus.WithContext(ctx.Context())
	log.Info("Get all questions handler")

	result, err := h.useCase.FindAll(ctx.Context())
	if err != nil {
		log.Errorf("Error getting all questions: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	log.Info("Questions found")
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": result})
}

// @Summary Get question by ID
// @Description Retrieve details of a specific question by its ID
// @Tags Questions
// @Param id path uint true "Question ID"
// @Produce json
// @Success 200 {object} responses.QuestionResponse "Question details"
// @Failure 400 {object} map[string]interface{} "Invalid question ID"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /questions/{id} [get]
func (h *QuestionHandler) GetQuestionByID(ctx *fiber.Ctx) error {
	log := logrus.WithContext(ctx.Context())
	log.Info("Get question handler")

	id := ctx.Params("id")

	transformId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Errorf("Error parsing id: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid id"})
	}

	newId := uint(transformId)

	result, err := h.useCase.FindByID(ctx.Context(), newId)
	if err != nil {
		log.Error(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	log.Info("Question found")
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": result,
	})
}

// @Summary Create a new question
// @Description Add a new question to the system
// @Tags Questions
// @Accept json
// @Produce json
// @Param question body requests.CreateQuestionRequest true "Question details"
// @Success 201 {object} map[string]interface{} "Question created"
// @Failure 400 {object} map[string]interface{} "Invalid request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /questions [post]
func (h *QuestionHandler) CreateQuestion(ctx *fiber.Ctx) error {
	log := logrus.WithContext(ctx.Context())
	log.Info("Create question handler")

	var req requests.CreateQuestionRequest
	if err := ctx.BodyParser(&req); err != nil {
		log.Errorf("Error parsing request: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	err := h.useCase.CreateQuestion(ctx.Context(), &req)
	if err != nil {
		log.Error(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	log.Info("Question created")
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Question created"})
}

// @Summary Update a question
// @Description Update the details of an existing question
// @Tags Questions
// @Accept json
// @Produce json
// @Param id path uint true "Question ID"
// @Param question body requests.CreateQuestionRequest true "Updated question details"
// @Success 200 {object} map[string]interface{} "Question updated"
// @Failure 400 {object} map[string]interface{} "Invalid request or question ID"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /questions/{id} [put]
func (h *QuestionHandler) UpdateQuestion(ctx *fiber.Ctx) error {
	log := logrus.WithContext(ctx.Context())
	log.Info("Update question handler")

	id := ctx.Params("id")

	transformId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Errorf("Error parsing id: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid id"})
	}

	newId := uint(transformId)

	var req requests.CreateQuestionRequest
	if err := ctx.BodyParser(&req); err != nil {
		log.Errorf("Error parsing request: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	err = h.useCase.UpdateQuestion(ctx.Context(), &req, newId)
	if err != nil {
		log.Error(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	log.Info("Question updated")
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Question updated"})
}

// @Summary Delete a question
// @Description Remove a question from the system
// @Tags Questions
// @Param id path uint true "Question ID"
// @Produce json
// @Success 200 {object} map[string]interface{} "Question deleted"
// @Failure 400 {object} map[string]interface{} "Invalid question ID"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /questions/{id} [delete]
func (h *QuestionHandler) DeleteQuestion(ctx *fiber.Ctx) error {
	log := logrus.WithContext(ctx.Context())
	log.Info("Delete question handler")

	id := ctx.Params("id")

	transformId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Errorf("Error parsing id: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid id"})
	}

	newId := uint(transformId)

	err = h.useCase.DeleteQuestion(ctx.Context(), newId)
	if err != nil {
		log.Error(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	log.Info("Question deleted")
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Question deleted"})
}

// @Summary Full text search for questions
// @Description Search for questions using a text query
// @Tags Questions
// @Param search query string true "Search query"
// @Produce json
// @Success 200 {object} responses.QuestionResponse "Search results"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /questions/search [get]
func (h *QuestionHandler) FullTextSearch(ctx *fiber.Ctx) error {
	log := logrus.WithContext(ctx.Context())
	log.Info("Full text search handler")

	search := ctx.Query("search")

	result, err := h.useCase.FullTextSearch(ctx.Context(), search)
	if err != nil {
		log.Error(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	log.Info("Questions found")
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": result})
}
