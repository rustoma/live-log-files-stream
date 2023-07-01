package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

var (
	logger              zerolog.Logger
	isLoggerInitialized bool
)

func InitLogger() *os.File {
	file, err := os.OpenFile(
		"myapp.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	if err != nil {
		panic(err)
	}

	if isProdDev() {
		logger = zerolog.New(file).With().Timestamp().Logger()
	} else {
		logger = log.Logger
	}

	isLoggerInitialized = true

	return file
}
