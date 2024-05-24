package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

// User là struct đại diện cho một người dùng
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// GetAllUsers lấy tất cả người dùng từ cơ sở dữ liệu
func GetAllUsers() ([]User, error) {
	// Get database reference
	db := Database

	// Define a filter to get all users
	filter := bson.M{}

	// Find all users
	cursor, err := db.Collection("users").Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var users []User
	// Iterate through the cursor and decode each user
	for cursor.Next(context.Background()) {
		var user User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
