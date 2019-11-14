package db

import (
	"context"

	"github.com/jinzhu/gorm"
)

// Muscle struct defines a muscle
type Muscle struct {
	gorm.Model
	Name        string
	Description string
}

// Repository interface provides access to the muscles storage
type Repository interface {
	ListMuscles(ctx context.Context) ([]Muscle, error)
}

////////////////////////////////////////

// MuscleStorage struct stores muscle data in DB
type MuscleStorage struct {
	db *gorm.DB
}

// NewMuscleStorage func returns a new muscles storage
func NewMuscleStorage() *MuscleStorage {
	storage := new(MuscleStorage)
	storage.db = OpenConnection()
	storage.db.AutoMigrate(&Muscle{})

	return storage
}

// ListMuscles func returns all muscles
func (s *MuscleStorage) ListMuscles(ctx context.Context) ([]Muscle, error) {
	var muscles []Muscle
	err := s.db.Find(&muscles).Error
	return muscles, err
}
