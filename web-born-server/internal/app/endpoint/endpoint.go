package endpoint

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Service interface {
	DaysLeft() int16
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) StatusHandler(ctx echo.Context) error {
	err := ctx.String(http.StatusOK, "Server active")
	if err != nil {
		return err
	}
	return nil
}

func (e *Endpoint) DateBornHandler(ctx echo.Context) error {
	dur := e.s.DaysLeft()
	res_s := fmt.Sprintf("Days until born: %d", dur)
	err := ctx.String(http.StatusOK, res_s)
	if err != nil {
		return err
	}
	return nil
}
