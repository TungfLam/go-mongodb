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

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	objReturn.Msg = `lấy dữ liệu thành công`
	objReturn.Data = users

	json.NewEncoder(w).Encode(objReturn)
}
