package repositories

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDB struct {
	mongo *mongo.Database
}

type UserRepo interface {
	Create(data interface{}) error
}

func NewUserRepo(db *mongo.Database) UserRepo {
	return &MongoDB{
		mongo: db,
	}
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
