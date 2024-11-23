package main

import (
	"log"
	"talana_prueba_tecnica/src/app/module"
	"talana_prueba_tecnica/src/shared"

	_ "talana_prueba_tecnica/docs"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title Talana prueba tecnica
// @version 1.0
// @description API para la gestión de trivias, destinada a ser prueba tecnica de juan martinez simi, para talana
// @host localhost:8080
// @BasePath /

func main() {
	log.Println("Starting server...")
	e := fiber.New()
	envs := shared.GetEnvs()
	shared.Init()

	e.Get("/swagger/*", fiberSwagger.WrapHandler)

	module.UserModule(e)
	module.QuestionModule(e)
	module.TriviaModule(e)

	err := e.Listen(":" + envs["PORT"])
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
