package apis

import (
	"net/http"

	"github.com/allenjoseph/workout-metrics-api/internal/db"
	"github.com/allenjoseph/workout-metrics-api/internal/services"
	"github.com/allenjoseph/workout-metrics-api/internal/util"
	"github.com/gorilla/mux"
)

var service services.MuscleService

// RegisterMusclesRoutes function
func RegisterMusclesRoutes(r *mux.Router) {
	router := r.PathPrefix("muscles").Subrouter()

	router.HandleFunc("/", getMuscles).Methods("GET")
	router.HandleFunc("/{id}", getMuscle).Methods("GET")

	client := db.OpenConnection()
	repository := db.MuscleRepositoryImpl{
		Client: client,
	}
	service = services.MuscleService{
		Repo: &repository,
	}
}

func getMuscles(w http.ResponseWriter, r *http.Request) {
	muscles, err := service.ListMuscles(r.Context())
	if err != nil {
		util.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	util.RespondJSON(w, http.StatusOK, muscles)
}

func getMuscle(w http.ResponseWriter, r *http.Request) {
	//..
}
