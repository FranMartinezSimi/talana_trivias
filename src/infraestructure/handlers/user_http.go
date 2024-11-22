package handlers

import (
	"talana_prueba_tecnica/src/app/usecases"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	usecase usecases.UserUseCase
}

func NewUserHandler(usecase usecases.UserUseCase) UserHandler {
	return UserHandler{usecase: usecase}
}

func (h *UserHandler) GetAllUsers(ctx *fiber.Ctx) error {
	log := logrus.WithContext(ctx.Context())
	log.Info("Get all users handler")

	result, err := h.usecase.FindAll(ctx.Context())
	if err != nil {
		log.Error(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	log.Info("Users found")
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": result})
}
