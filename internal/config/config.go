package config

import (
	"os"

	"github.com/joho/godotenv"
)

// AuthProvider is a general object to store oauth2 client information
type AuthProvider struct {
	ClientID     string
	ClientSecret string
}

// Firefly stores auth information for the Firefly III oauth client
type Firefly struct {
	AuthProvider
}

// Monzo stores auth information for the Monzo oauth client
type Monzo struct {
	AuthProvider
}

// Config stores data for the running application
type Config struct {
	AppRoot   string
	MongoURL  string
	Firefly   Firefly
	Monzo     Monzo
	StartDate string
}

// New creates a new Config object
func New() *Config {
	godotenv.Load()

	return &Config{
		AppRoot:  os.Getenv("APP_ROOT"),
		MongoURL: os.Getenv("MONGO_URL"),
		Monzo: Monzo{
			AuthProvider: AuthProvider{
				ClientID:     os.Getenv("MONZO_CLIENT_ID"),
				ClientSecret: os.Getenv("MONZO_CLIENT_SECRET"),
			},
		},
		Firefly: Firefly{
			AuthProvider: AuthProvider{
				ClientID:     os.Getenv("FIREFLY_CLIENT_ID"),
				ClientSecret: os.Getenv("FIREFLY_CLIENT_SECRET"),
			},
		},
	}
}
