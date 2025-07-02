package app

import (
	"log"

	"example.com/rest_full_example/internal/app/endpoint"
	"example.com/rest_full_example/internal/app/middleware"
	"example.com/rest_full_example/internal/app/service"
	"github.com/labstack/echo/v4"
)

type App struct {
	e    *endpoint.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}
	a.s = service.New()

	a.e = endpoint.New(a.s)

	a.echo = echo.New()

	a.echo.Use(middleware.RoleCheck)

	a.echo.GET("/status", a.e.StatusHandler)
	a.echo.GET("/born", a.e.DateBornHandler)
	return a, nil
}

func (a *App) Run() error {
	log.Println("Server starting")
	err := a.echo.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
