package services

import (
	"context"

	"github.com/allenjoseph/workout-metrics-api/internal/db"
)

// Service interface provides list muscles operations
type Service interface {
	ListMuscles(ctx context.Context) ([]db.Muscle, error)
	CreateMuscle(ctx context.Context, muscle *db.Muscle)
}

type service struct {
	repo *db.MuscleStorage
}

// NewService func creates an muscle service with the necessary dependencies
func NewService(repo *db.MuscleStorage) Service {
	return &service{repo}
}

// ListMuscles func list of muscles from the database
func (s *service) ListMuscles(ctx context.Context) ([]db.Muscle, error) {
	return s.repo.ListMuscles(ctx)
}

// CreateMuscle func list of muscles from the database
func (s *service) CreateMuscle(ctx context.Context, muscle *db.Muscle) {
	s.repo.CreateMuscle(ctx, muscle)
}
