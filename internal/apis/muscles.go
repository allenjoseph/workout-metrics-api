package apis

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterMusclesRoutes(r *mux.Router) {
	router := r.PathPrefix("muscles").Subrouter()

	router.HandleFunc("/", getMuscles).Methods("GET")
	router.HandleFunc("/{id}", getMuscle).Methods("GET")
}

func getMuscles(w http.ResponseWriter, r *http.Request) {
	//..
}

func getMuscle(w http.ResponseWriter, r *http.Request) {
	//..
}
