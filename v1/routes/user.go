package routes

import (
	"go-dashboard/util/jwt"
	"go-dashboard/v1/controllers"
	"go-dashboard/v1/repositories"
	"go-dashboard/v1/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

func User(e *echo.Echo, db *mongo.Database) *echo.Echo {
	var (
		userRepo repositories.UserRepo       = repositories.NewUserRepo(db)
		userSvc  services.UserServices       = services.NewUserServices(userRepo)
		userCtrl controllers.UserControllers = controllers.NewUserControllers(userSvc)
	)

	v1 := e.Group("/v1/")
	v1.POST("login", userCtrl.Login)
	v1.Use(middleware.JWTWithConfig(jwt.JwtConfig()))
	v1.POST("user/create", userCtrl.Create)

	return e

}
