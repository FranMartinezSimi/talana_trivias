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
