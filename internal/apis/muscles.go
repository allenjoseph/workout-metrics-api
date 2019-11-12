package apis

import (
	"net/http"

	"github.com/allenjoseph/workout-metrics-api/internal/db"
	"github.com/allenjoseph/workout-metrics-api/internal/services"
	"github.com/allenjoseph/workout-metrics-api/internal/util"
	"github.com/gorilla/mux"
)

// RegisterMusclesRoutes function
func RegisterMusclesRoutes(r *mux.Router) {

	musclesStorage := db.NewMuscleStorage()
	musclesService := services.NewService(musclesStorage)

	sr := r.PathPrefix("/muscles").Subrouter()

	sr.HandleFunc("/", getMuscles(musclesService)).Methods("GET")
	sr.HandleFunc("/{id}/", getMuscle).Methods("GET")
}

func getMuscles(s services.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		muscles, err := s.ListMuscles(r.Context())
		if err != nil {
			util.RespondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		util.RespondJSON(w, http.StatusOK, muscles)
	}
}

func getMuscle(w http.ResponseWriter, r *http.Request) {
	//..
}
