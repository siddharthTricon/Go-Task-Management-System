package utils

import (
	"log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	
)

// LoadEnv loads environment variables from a .env file
func LoadEnv() {
	// Specify the file name of the configuration file
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".") // Look for the file in the current directory

	// Read in the configuration file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading .env file: %v", err)
	}

	// Enable automatic environment variable overrides
	viper.AutomaticEnv()

	log.Println("Environment variables loaded successfully")
}

// GetEnv retrieves an environment variable's value
func GetEnv(key string) string {
	value := viper.GetString(key)
	if value == "" {
		log.Printf("Warning: environment variable '%s' not set", key)
	}
	return value
}

func ValidateToken(tokenString string) bool {
    // Parse the JWT token
    token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
        return []byte(GetEnv("JWT_SECRET")), nil
    })

    if err != nil || !token.Valid {
        log.Println("Invalid token:", err)
        return false
    }
    return true
}


