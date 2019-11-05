package apis

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterExercisesRoutes(r *mux.Router) {
	router := r.PathPrefix("exercises").Subrouter()

	router.HandleFunc("/", getExercises).Methods("GET")
}

func getExercises(w http.ResponseWriter, r *http.Request) {
	//..
}
