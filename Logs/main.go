package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{}) // Enable JSON format
	log.SetLevel(logrus.WarnLevel)            // Only logs Warn, Error, and Fatal

	// Simple log message
	log.Info("Application started")

	// Log with fields wiht structured logging
	log.WithFields(logrus.Fields{
		"user": "Utkarsh",
		"role": "admin",
	}).Info("User logged in")

	// Different log levels
	log.Warn("This is a warning message")
	log.Error("This is an error message")
}
