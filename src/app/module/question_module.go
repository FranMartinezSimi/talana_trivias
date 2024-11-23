package module

import (
	questionsusecase "talana_prueba_tecnica/src/app/usecases/questions_usecase"
	"talana_prueba_tecnica/src/infraestructure/handlers"
	questionsrepository "talana_prueba_tecnica/src/infraestructure/repository/questions_repository"
	"talana_prueba_tecnica/src/shared"

	"github.com/gofiber/fiber/v2"
)

func QuestionModule(app *fiber.App) {
	db := shared.Init()
	questionRepo := questionsrepository.NewQuestionRepository(db)
	useCase := questionsusecase.NewQuestionsUseCase(questionRepo)
	handler := handlers.NewQuestionHandler(useCase)

	app.Get("/questions", handler.GetAllQuestions)
	app.Get("/questions/:id", handler.GetQuestionByID)
	app.Get("/questions ", handler.FullTextSearch)
	app.Post("/questions", handler.CreateQuestion)
	app.Put("/questions/:id", handler.UpdateQuestion)
	app.Delete("/questions/:id", handler.DeleteQuestion)
}
