package apis

import (
	"encoding/json"
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
	sr.HandleFunc("/", addMuscle(musclesService)).Methods("POST")
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

func addMuscle(s services.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var muscle db.Muscle
		err := decoder.Decode(&muscle)
		util.LogOnError("Failed to add muscle", err)
		s.CreateMuscle(r.Context(), &muscle)

		util.RespondJSON(w, http.StatusOK, muscle)
	}
}

func getMuscle(w http.ResponseWriter, r *http.Request) {
	//..
}
