package v1

import (
	"go-dashboard/util"
	"go-dashboard/v1/middlewares"
	"go-dashboard/v1/routes"
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
	e.Use(middlewares.Cors)
	routes.User(e, db)
	util.ViewRoutes(e.Routes())
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))

}
