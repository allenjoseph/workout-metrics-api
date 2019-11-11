package db

import (
	"github.com/allenjoseph/workout-metrics-api/internal/util"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

// Client struct
type Client struct {
	db *gorm.DB
}

// OpenConnection function
func OpenConnection() *Client {
	dataSource := viper.GetString("db.datasource")
	db, err := gorm.Open("postgres", dataSource)
	defer db.Close()

	util.FailOnError("Failed to open db connection", err)
	util.LogInfo("DB connection opened")
	return &Client{db}
}
