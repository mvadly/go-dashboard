package routes

import (
	"go-dashboard/v1/controllers"
	"go-dashboard/v1/repositories"
	"go-dashboard/v1/services"
	"net/http"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
)

func User(e *echo.Echo, db *mongo.Database) *echo.Echo {
	var (
		userRepo repositories.UserRepo       = repositories.NewUserRepo(db)
		userSvc  services.UserServices       = services.NewUserServices(userRepo)
		userCtrl controllers.UserControllers = controllers.NewUserControllers(userSvc)
	)

	v1 := e.Group("v1/user/")
	v1.GET("hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	v1.POST("create", userCtrl.Create)

	return e

}
