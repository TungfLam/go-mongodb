package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetAllUsers() ([]User, error) {
	db := Database

	filter := bson.M{}

	cursor, err := db.Collection("users").Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var users []User
	for cursor.Next(context.Background()) {
		var user User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
