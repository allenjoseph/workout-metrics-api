package db

import (
	"context"

	"github.com/jinzhu/gorm"
)

// Muscle struct defines a muscle
type Muscle struct {
	CommonModelFields
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Repository interface provides access to the muscles storage
type Repository interface {
	ListMuscles(ctx context.Context) ([]Muscle, error)
	CreateMuscle(ctx context.Context, muscle *Muscle)
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

// CreateMuscle func create a muscle record
func (s *MuscleStorage) CreateMuscle(ctx context.Context, muscle *Muscle) {
	s.db.Create(&muscle)
}
