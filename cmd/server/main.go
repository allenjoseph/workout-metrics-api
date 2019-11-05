package main

import (
	"net/http"

	"github.com/allenjoseph/workout-metrics-api/internal/apis"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	apis.RegisterMusclesRoutes(r)
	apis.RegisterExercisesRoutes(r)

	http.ListenAndServe(":8080", r)
}
