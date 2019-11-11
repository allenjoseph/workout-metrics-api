package util

import (
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

// FailOnError util function
func FailOnError(msg string, err error) {
	if logger == nil {
		logger = logrus.New()
	}
	if err != nil {
		logger.Panic(msg, err)
	}
}

// LogInfo util function
func LogInfo(msg string) {
	logger.Info(msg)
}
