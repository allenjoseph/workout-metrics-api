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

// LogOnError util function
func LogOnError(msg string, err error) {
	if logger == nil {
		logger = logrus.New()
	}
	if err != nil {
		LogInfo(msg, err)
	}
}

// LogInfo util function
func LogInfo(msg string, extra interface{}) {
	logger.Info(msg, extra)
}
