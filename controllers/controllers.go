// controllers/controller.go

package controllers

import (
	"encoding/json"
	"net/http"
	"study_learn_go/models"
)

var objReturn = struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data,omitempty"`
}{
	Status: 1,
	Msg:    "OK",
}

// GetAllUsersHandler lấy tất cả người dùng
func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Call the function to get all users
	users, err := models.GetAllUsers()
	if err != nil {
		// Handle the error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the users as JSON response
	json.NewEncoder(w).Encode(users)
}
