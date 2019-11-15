package db

import (
	"time"

	"github.com/allenjoseph/workout-metrics-api/internal/util"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

// CommonModelFields struct defines the default gorm.Model and its json representation
type CommonModelFields struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// OpenConnection function
func OpenConnection() *gorm.DB {
	dataSource := viper.GetString("db.datasource")
	db, err := gorm.Open("postgres", dataSource)

	util.FailOnError("Failed to open db connection", err)
	util.LogInfo("DB connection opened")
	return db
}
