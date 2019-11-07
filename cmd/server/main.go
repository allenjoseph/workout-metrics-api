package main

import (
	"net/http"

	"github.com/allenjoseph/workout-metrics-api/internal/apis"
	"github.com/allenjoseph/workout-metrics-api/internal/config"
	"github.com/allenjoseph/workout-metrics-api/internal/db"
)

func main() {
	config.LoadConfig()
	db.OpenConnection()

	appRoutes := apis.RegisterRoutes()

	panic(http.ListenAndServe(":8080", appRoutes))
}
