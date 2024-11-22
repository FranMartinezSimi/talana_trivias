package module

import (
	"talana_prueba_tecnica/src/app/usecases"
	"talana_prueba_tecnica/src/infraestructure/handlers"
	"talana_prueba_tecnica/src/infraestructure/repository"
	"talana_prueba_tecnica/src/shared"

	"github.com/gofiber/fiber/v2"
)

func UserModule(app *fiber.App) {
	db := shared.Init()
	userRepo := repository.NewUserRepository(db)
	userUseCase := usecases.NewUserUseCase(userRepo)
	userHandler := handlers.NewUserHandler(*userUseCase)

	app.Get("/users", userHandler.GetAllUsers)
	app.Get("/users/:id", userHandler.GetUserByID)
	app.Post("/users", userHandler.CreateUser)
	app.Put("/users/:id", userHandler.UpdateUser)
	app.Delete("/users/:id", userHandler.DeleteUser)
}
