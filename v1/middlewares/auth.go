package middlewares

import (
	"os"

	"github.com/labstack/echo/v4"
)

func Cors(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		clientUrl := os.Getenv("CLIENT_DOMAIN")
		c.Request().Header.Add("Access-Control-Allow-Origin", clientUrl)
		c.Request().Header.Add("Access-Control-Allow-Credentials", "true")
		c.Request().Header.Add("Access-Control-Allow-Headers", "token, Content-Type, Content-Length, Accept-Encoding, X-XSRF-TOKEN, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Request().Header.Add("Access-Control-Allow-Methods", "GET, POST, HEAD, PATCH, PUT, DELETE, OPTIONS")
		return next(c)

	}
}
