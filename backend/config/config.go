package config

import (
	"os"
)

func GetCredentialsPath() string {
	credentialsPath := os.Getenv("GOOGLE_CREDENTIALS_PATH")
	if credentialsPath == "" {
		credentialsPath = "../../credentials.json"
	}
	return credentialsPath
}
