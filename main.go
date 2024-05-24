package main

import (
	"fmt"
	"net/http"
	"study_learn_go/models"
	"study_learn_go/routes"
)

func main() {
	r := routes.SetupRoutes()

	err := models.InitDB()
	if err != nil {
		fmt.Println("Failed to initialize database:", err)
		return
	}

	fmt.Println("Starting server on :3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
