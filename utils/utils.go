package utils

import (
	"os"
	"strconv"
)

func GetPort() string {
	// Try to get the PORT environment variable
	envPort := os.Getenv("PORT")
	if envPort != "" {
		// Check if the environment variable is a valid port number
		port, err := strconv.Atoi(envPort)
		if err == nil && port >= 1 && port <= 65535 {
			return envPort
		}
	}

	// Default to port 3000 if the environment variable is not set or invalid
	return "3000"
}
