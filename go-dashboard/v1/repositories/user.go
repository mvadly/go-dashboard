package repositories

import (
	"context"
	"fmt"
	"go-dashboard/v1/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDB struct {
	mongo *mongo.Database
}

type UserRepo interface {
	Login(data interface{}) (results models.Users, err error)
	Create(data interface{}) error
}

func NewUserRepo(db *mongo.Database) UserRepo {
	return &MongoDB{
		mongo: db,
	}
}

func (db *MongoDB) Login(data interface{}) (results models.Users, err error) {
	ctx := context.TODO()
	login := db.mongo.Collection("users").FindOne(ctx, data)
	if login.Err() != nil {
		fmt.Println("disini errornya")
		return results, login.Err()
	}

	if err = login.Decode(&results); err != nil {
		fmt.Println("disini errornya2")
		return results, err
	}

	return results, nil

}

func (db *MongoDB) Create(data interface{}) error {
	ctx := context.TODO()
	_, err := db.mongo.Collection("users").InsertOne(ctx, data)
	if err != nil {
		fmt.Println("error mongo connection: ", err)
		return err
	}

	return nil
}
