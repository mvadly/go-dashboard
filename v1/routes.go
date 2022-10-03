package v1

import (
	"go-dashboard/v1/middlewares"
	"go-dashboard/v1/routes"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

func EchoRoute(db *mongo.Database) {

	e := echo.New()
	e.Debug = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middlewares.Cors)
	routes.User(e, db)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))

}
