package v1

import (
	"go-dashboard/util"
	"go-dashboard/v1/routes"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

func EchoRoute(db *mongo.Database) {

	e := echo.New()
	e.Debug = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	// e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
	// 	TokenLookup: "header:X-XSRF-TOKEN",
	// }))
	routes.User(e, db)
	e.HTTPErrorHandler = customErrorHandler
	util.ViewRoutes(e.Routes())
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))

}

func customErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	msg := "unknown error"
	he, ok := err.(*echo.HTTPError)
	if ok {
		code = he.Code
		msg = he.Message.(string)
	}

	util.JSON(c, code, util.ResJSON{
		Code:    "01",
		Message: msg,
		Data:    nil,
	})

}
