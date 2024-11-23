package module

import (
	"github.com/gofiber/fiber/v2"
	"talana_prueba_tecnica/src/app/usecases/game_usecase"
	"talana_prueba_tecnica/src/infraestructure/handlers"
	"talana_prueba_tecnica/src/infraestructure/repository/game_repository"
	questionsrepository "talana_prueba_tecnica/src/infraestructure/repository/questions_repository"
	triviarepository "talana_prueba_tecnica/src/infraestructure/repository/trivia_repository"
	"talana_prueba_tecnica/src/shared"
)

func GameModule(app *fiber.App) {
	db := shared.Init()
	gameRepo := game_repository.NewGameRepository(db)
	questionRepo := questionsrepository.NewQuestionRepository(db)
	triviaReRepo := triviarepository.NewTriviaRepository(db)
	gameUseCase := game_usecase.NewGameUseCase(gameRepo, questionRepo, triviaReRepo)
	gamerHandler := handlers.NewGameHandler(gameUseCase)

	app.Get("/games/trivias/:id/questions", gamerHandler.GetQuestionsForTrivia)
	app.Post("/games/trivias/:id/answers", gamerHandler.SubmitAnswers)

}
