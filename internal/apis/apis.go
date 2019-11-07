package apis

import (
	"github.com/gorilla/mux"
)

// RegisterRoutes function
func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()

	RegisterMusclesRoutes(router)
	RegisterExercisesRoutes(router)

	return router
}
