package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logger.Panic().Err(err).Msg("")
	}

	logFile := InitLogger()
	defer logFile.Close()

	if !isLoggerInitialized {
		panic("Logger needs to be initialized")
	}

	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)

	mux.Get("/api/v1", makeHTTPHandler(Home))

	logger.Info().Msg("Server is running on port :8080")
	err = http.ListenAndServe(":8080", mux)

	if err != nil {
		logger.Panic().Err(err).Msg("")
	}
}
