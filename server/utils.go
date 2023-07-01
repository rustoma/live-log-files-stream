package main

import (
	"os"
)

func isProdDev() bool {
	if os.Getenv("APP_ENV") == "prod" {
		return true
	}
	return false
}
