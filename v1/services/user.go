package services

import (
	"go-dashboard/util/encryption"
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
	err = s.repo.Create(user)
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = "internal server error"
		res.Response = err.Error()
		return res, err
	}

	if user.CreatedBy == "" {
		user.CreatedBy = "system"
	}

	user.Password = rand
	res.Code = http.StatusOK
	res.Message = "success create user"
	res.Response = user
	return res, nil
}
