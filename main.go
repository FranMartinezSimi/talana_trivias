package main

import (
	"log"
	"talana_prueba_tecnica/src/app/module"
	"talana_prueba_tecnica/src/shared"

	"github.com/gofiber/fiber/v2"
)

func main() {
	log.Println("Starting server...")
	e := fiber.New()
	envs := shared.GetEnvs()
	shared.Init()

	module.UserModule(e)
	module.QuestionModule(e)

	err := e.Listen(":" + envs["PORT"])
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
