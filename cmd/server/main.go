package main

import (
	"net/http"

	"github.com/allenjoseph/workout-metrics-api/internal/apis"
	"github.com/allenjoseph/workout-metrics-api/internal/config"
)

func main() {
	config.LoadConfig()

	appRoutes := apis.RegisterRoutes()

	panic(http.ListenAndServe(":8080", appRoutes))
}
