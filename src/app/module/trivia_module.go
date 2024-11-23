package module

import (
	triviausecase "talana_prueba_tecnica/src/app/usecases/trivia_usecase"
	"talana_prueba_tecnica/src/infraestructure/handlers"
	triviarepository "talana_prueba_tecnica/src/infraestructure/repository/trivia_repository"
	repository "talana_prueba_tecnica/src/infraestructure/repository/user_repository"
	"talana_prueba_tecnica/src/shared"

	"github.com/gofiber/fiber/v2"
)

func TriviaModule(app *fiber.App) {
	db := shared.Init()
	triviaRepo := triviarepository.NewTriviaRepository(db)
	userRepo := repository.NewUserRepository(db)
	triviaUseCase := triviausecase.NewTriviaUseCase(triviaRepo, userRepo)
	triviaHandler := handlers.NewTriviaHandler(triviaUseCase)

	app.Get("/trivias", triviaHandler.GetAllTrivias)
	app.Get("/trivias/:id", triviaHandler.GetTriviaByID)
	app.Post("/trivias", triviaHandler.CreateTrivia)
	app.Put("/trivias/:id", triviaHandler.UpdateTrivia)
	app.Delete("/trivias/:id", triviaHandler.DeleteTrivia)

}
