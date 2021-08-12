package checkEnv

import (
	"log"
	"os"
)

func CheckForRequiredEnvVars(l *log.Logger, varNames []string) bool {
	isValid := true

	for _, key := range varNames {
		if os.Getenv(key) == "" {
			l.Printf("[ERROR]: Environment variable '%s' not set.\n", key)
			isValid = false
		}
	}

	return isValid
}
