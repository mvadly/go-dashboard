package controllers

import (
	"go-dashboard/util"
	"go-dashboard/v1/models"
	"go-dashboard/v1/services"
	"net/http"

	"github.com/labstack/echo"
)

type userCtrl struct {
	svc services.UserServices
}

type UserControllers interface {
	Create(c echo.Context) error
}

func NewUserControllers(svc services.UserServices) UserControllers {
	return &userCtrl{
		svc: svc,
	}
}

func (con *userCtrl) Create(c echo.Context) error {
	var user models.Users
	err := c.Bind(&user)
	if err != nil {
		util.JSON(c, http.StatusBadRequest, util.ResJSON{
			Code:    "01",
			Message: err.Error(),
		})
	}

	res, err := con.svc.Create(user)
	if err != nil {
		util.JSON(c, res.Code, util.ResJSON{
			Code:    "01",
			Message: err.Error(),
		})
	}

	util.JSON(c, http.StatusOK, util.ResJSON{
		Code:    "00",
		Message: res.Message,
		Data:    res.Response,
	})

	return nil
}
