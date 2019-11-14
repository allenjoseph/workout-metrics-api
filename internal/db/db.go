package db

import (
	"github.com/allenjoseph/workout-metrics-api/internal/util"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

// OpenConnection function
func OpenConnection() *gorm.DB {
	dataSource := viper.GetString("db.datasource")
	db, err := gorm.Open("postgres", dataSource)

	util.FailOnError("Failed to open db connection", err)
	util.LogInfo("DB connection opened", nil)
	return db
}
