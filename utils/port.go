package utils

import (
	"os"
	"strconv"
)

func GetPort() string {
	envPort := os.Getenv("PORT")
	if envPort != "" {
		port, err := strconv.Atoi(envPort)
		if err == nil && port >= 1 && port <= 65535 {
			return envPort
		}
	}
	return "3000"
}
