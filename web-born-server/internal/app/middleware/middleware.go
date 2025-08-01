package middleware

import (
	"log"
	"strings"

	"github.com/labstack/echo/v4"
)

const roleAdmin = "admin"

func RoleCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		val := ctx.Request().Header.Get("User-Role")
		if strings.EqualFold(val, roleAdmin) {
			log.Println("Admin in space!")
		}

		err := next(ctx)
		if err != nil {
			return err
		}
		return nil
	}
}
