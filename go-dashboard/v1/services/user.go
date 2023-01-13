package services

import (
	"errors"
	"go-dashboard/util"
	"go-dashboard/util/encryption"
	"go-dashboard/util/jwt"
	"go-dashboard/v1/models"
	"go-dashboard/v1/repositories"
	"net/http"
)

type UserServices struct {
	repo repositories.UserRepo
}

func NewUserServices(repo repositories.UserRepo) UserServices {
	return UserServices{
		repo: repo,
	}
}

func (s *UserServices) Login(user models.Users) (res models.ReponseServices, err error) {
	user.Password = user.Username + user.Password
	data, err := s.repo.Login(map[string]interface{}{"username": user.Username})
	if err != nil || !encryption.CheckPasswordHash(user.Password, data.Password) {
		res.Code = http.StatusNotFound
		res.Message = "username and password not match"
		res.Response = nil
		return res, errors.New(res.Message)
	}

	token, err := jwt.SetToken(data)
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = "create token failed"
		res.Response = err.Error()
		return res, err
	}

	res.Code = http.StatusOK
	res.Message = "success"
	res.Response = map[string]interface{}{
		"access_token": token,
	}
	return res, err
}

func (s *UserServices) Create(user models.Users) (res models.ReponseServices, err error) {
	rand := encryption.GeneratePassword()
	pass, err := encryption.HashPassword(user.Username + rand)
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = "internal server error"
		res.Response = err.Error()
		return res, err
	}

	user.Password = pass
	user.CreatedAt = util.TimeNow()
	user.SetDefaultID()

	if user.CreatedBy == "" {
		user.CreatedBy = "system"
	}

	err = s.repo.Create(user)
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = "internal server error"
		res.Response = err.Error()
		return res, err
	}

	user.Password = rand
	res.Code = http.StatusOK
	res.Message = "success create user"
	res.Response = user
	return res, nil
}
