package main

import (
	"github.com/labstack/echo/v4"
	"talana_prueba_tecnica/src/shared"
)

func main() {
	e := echo.New()
	envs := shared.GetEnvs()
	shared.Init()

	err := e.Start(":" + envs["PORT"])
	if err != nil {
		e.Logger.Fatal(err)
	}
}
