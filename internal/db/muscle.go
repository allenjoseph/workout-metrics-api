package db

import (
	"context"

	"github.com/jinzhu/gorm"
)

// Muscle struct
type Muscle struct {
	gorm.Model
	Name        string
	Description string
}

// MuscleRepository interface
type MuscleRepository interface {
	ListMuscles(ctx context.Context) ([]Muscle, error)
	Close()
}

var r MuscleRepository

// SetRepository func
func SetRepository(repository MuscleRepository) {
	r = repository
}

// ListMuscles func
func ListMuscles(ctx context.Context) ([]Muscle, error) {
	return r.ListMuscles(ctx)
}

// Close func
func Close() {
	r.Close()
}

// MuscleRepositoryImpl struct
type MuscleRepositoryImpl struct {
	Client *Client
}

// ListMuscles implementation
func (r *MuscleRepositoryImpl) ListMuscles(ctx context.Context) ([]Muscle, error) {
	var muscles []Muscle
	err := r.Client.db.Find(&muscles).Error
	return muscles, err
}
