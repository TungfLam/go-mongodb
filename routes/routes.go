package routes

import (
	"log"
	"net/http"
	"study_learn_go/controllers"

	"github.com/gorilla/mux"
)

// Set up routes
func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/", LogRequest(controllers.GetAllUsersHandler)).Methods("GET")
	// r.HandleFunc("/users", controllers.GetUsersHandler).Methods("GET")
	// r.HandleFunc("/users/{id}", controllers.GetUserByIDHandler).Methods("GET")
	// r.HandleFunc("/users", controllers.CreateUserHandler).Methods("POST")
	// r.HandleFunc("/users/{id}", controllers.UpdateUserHandler).Methods("PUT")
	// r.HandleFunc("/users/{id}", controllers.DeleteUserHandler).Methods("DELETE")

	return r
}
func LogRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw := &responseWriter{w, http.StatusOK}

		next.ServeHTTP(rw, r)

		log.Printf("%s %s %s %d", r.RemoteAddr, r.Method, r.URL, rw.status)
	}
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}
