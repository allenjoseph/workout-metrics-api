package services

import (
	"context"

	"github.com/allenjoseph/workout-metrics-api/internal/db"
)

// MuscleService struct
type MuscleService struct {
	Repo *db.MuscleRepositoryImpl
}

// ListMuscles func
func (service *MuscleService) ListMuscles(ctx context.Context) ([]db.Muscle, error) {
	return service.Repo.ListMuscles(ctx)
}
