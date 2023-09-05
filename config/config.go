package config

import (
  "os"

  "github.com/joho/godotenv"
)

// LoadEnv function to load environment variables from .env file
func LoadEnv() {
  // Load .env file
  err := godotenv.Load()
  if err != nil {
    panic("Error loading .env file")
  }
}

// GetEnv function to get environment variable by key
func GetEnv(key string) string {
  return os.Getenv(key)
}
