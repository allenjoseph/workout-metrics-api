package util

import (
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

// FailOnError util function
func FailOnError(msg string, err error) {
	if err != nil {
		logger.Panic(msg, err)
	}
}

// LogOnError util function
func LogOnError(msg string, err error) {
	if err != nil {
		LogInfo(msg, err)
	}
}

// LogInfo util function
func LogInfo(args ...interface{}) {
	logger.Infoln(args...)
}
