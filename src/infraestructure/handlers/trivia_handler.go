package handlers

import (
	triviausecase "talana_prueba_tecnica/src/app/usecases/trivia_usecase"
	"talana_prueba_tecnica/src/entity/requests"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type TriviaHandler struct {
	useCase triviausecase.TriviaUseCaseInterface
}

func NewTriviaHandler(useCase triviausecase.TriviaUseCaseInterface) *TriviaHandler {
	return &TriviaHandler{
		useCase: useCase,
	}
}

func (h *TriviaHandler) GetAllTrivias(ctx *fiber.Ctx) error {
	log := logrus.WithContext(ctx.Context())
	log.Info("Get all trivias handler")

	result, err := h.useCase.FindAll(ctx.Context())
	if err != nil {
		log.Errorf("Error getting all trivias: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	log.Info("Trivias found")
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": result})
}

func (h *TriviaHandler) GetTriviaByID(ctx *fiber.Ctx) error {
	log := logrus.WithContext(ctx.Context())
	log.Info("Get trivia by ID handler")

	id, err := ctx.ParamsInt("id")
	if err != nil {
		log.Errorf("Invalid trivia ID: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid trivia ID"})
	}

	result, err := h.useCase.FindByID(ctx.Context(), uint(id))
	if err != nil {
		log.Errorf("Error getting trivia by ID: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	log.Info("Trivia found")
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": result})
}

func (h *TriviaHandler) CreateTrivia(ctx *fiber.Ctx) error {
	log := logrus.WithContext(ctx.Context())
	log.Info("Create trivia handler")

	var req requests.CreateTriviaRequest
	if err := ctx.BodyParser(&req); err != nil {
		log.Errorf("Error parsing request: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := h.useCase.CreateTrivia(ctx.Context(), &req); err != nil {
		log.Errorf("Error creating trivia: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	log.Info("Trivia created")
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Trivia created successfully"})
}

func (h *TriviaHandler) UpdateTrivia(ctx *fiber.Ctx) error {
	log := logrus.WithContext(ctx.Context())
	log.Info("Update trivia handler")

	id, err := ctx.ParamsInt("id")
	if err != nil {
		log.Errorf("Invalid trivia ID: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid trivia ID"})
	}

	var req requests.CreateTriviaRequest
	if err := ctx.BodyParser(&req); err != nil {
		log.Errorf("Error parsing request: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := h.useCase.UpdateTrivia(ctx.Context(), &req, uint(id)); err != nil {
		log.Errorf("Error updating trivia: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	log.Info("Trivia updated")
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Trivia updated successfully"})
}

func (h *TriviaHandler) DeleteTrivia(ctx *fiber.Ctx) error {
	log := logrus.WithContext(ctx.Context())
	log.Info("Delete trivia handler")

	id, err := ctx.ParamsInt("id")
	if err != nil {
		log.Errorf("Invalid trivia ID: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid trivia ID"})
	}

	if err := h.useCase.DeleteTrivia(ctx.Context(), uint(id)); err != nil {
		log.Errorf("Error deleting trivia: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	log.Info("Trivia deleted")
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Trivia deleted successfully"})
}
