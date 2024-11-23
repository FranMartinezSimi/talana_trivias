package handlers

import (
	"strconv"
	usecases "talana_prueba_tecnica/src/app/usecases/user_usecase"
	"talana_prueba_tecnica/src/entity/requests"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	usecase usecases.UserUseCaseInterface
}

func NewUserHandler(usecase usecases.UserUseCaseInterface) UserHandler {
	return UserHandler{usecase: usecase}
}

func (h *UserHandler) GetAllUsers(ctx *fiber.Ctx) error {
	log := logrus.WithContext(ctx.Context())
	log.Info("Get all users handler")

	result, err := h.usecase.FindAll(ctx.Context())
	if err != nil {
		log.Errorf("Error getting all users: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	log.Info("Users found")
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"data": result})
}

func (h *UserHandler) GetUserByID(ctx *fiber.Ctx) error {
	log := logrus.WithContext(ctx.Context())
	log.Info("Get user handler")

	id := ctx.Params("id")

	transformId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Errorf("Error parsing id: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid id"})
	}

	newId := uint(transformId)
	result, err := h.usecase.GetUserByID(ctx.Context(), newId)
	if err != nil {
		log.Error(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	log.Info("User found")
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": result,
	})
}

func (h *UserHandler) CreateUser(ctx *fiber.Ctx) error {
	log := logrus.WithContext(ctx.Context())
	log.Info("Create user handler")

	var userRequest requests.RegisterUserRequest
	if err := ctx.BodyParser(&userRequest); err != nil {
		log.Error(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.usecase.CreateUser(ctx.Context(), userRequest)
	if err != nil {
		log.Error(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	log.Info("User created")
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created",
	})
}

func (h *UserHandler) UpdateUser(ctx *fiber.Ctx) error {
	log := logrus.WithContext(ctx.Context())
	log.Info("Update user handler")

	id := ctx.Params("id")
	transformId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Errorf("Error parsing id: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid id"})
	}

	newId := uint(transformId)

	var userRequest requests.UpdateUserRequest

	if err := ctx.BodyParser(&userRequest); err != nil {
		log.Error(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err = h.usecase.UpdateUser(ctx.Context(), newId, userRequest)

	if err != nil {
		log.Error(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	log.Info("User updated")
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User updated",
	})
}

func (h *UserHandler) DeleteUser(ctx *fiber.Ctx) error {
	log := logrus.WithContext(ctx.Context())
	log.Info("Delete user handler")

	id := ctx.Params("id")
	transformId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Errorf("Error parsing id: %v", err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid id"})
	}

	newId := uint(transformId)

	err = h.usecase.DeleteUser(ctx.Context(), newId)

	if err != nil {
		log.Error(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	log.Info("User deleted")
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User deleted",
	})
}
