package services

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/HerlambangHaryo/go-crud-laravel/app/models"
	"golang.org/x/crypto/bcrypt"
)

	// Secret key for signing and verifying jwt token
		var secretKey = []byte("secret")

	// ValidateUser function to validate user input
		func ValidateUser(user *models.User) error {
			// Create a validator instance
			v := validator.New()

			// Validate user struct using validator tags
			err := v.Struct(user)
			if err != nil {
				return err
			}

			// Return nil if no error
			return nil
		}

	// HashPassword function to hash user password using bcrypt
		func HashPassword(password string) (string, error) {
			// Generate a bcrypt hash from the password
			hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if err != nil {
				return "", err
			}

			// Return the hash as a string and error if any
			return string(hash), nil
		}

	// CheckPassword function to check user password with hashed password using bcrypt
		func CheckPassword(password string, hash string) bool {
			// Compare the password and the hash using bcrypt
			err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

			// Return true if no error, false otherwise
			return err == nil
		}

	// GenerateToken function to generate jwt token for user using user id
		func GenerateToken(user *models.User) (string, error) {
			// Create a jwt token with user id as claim and expiration time of 24 hours
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"user_id": user.ID,
				"exp":     time.Now().Add(time.Hour * 24).Unix(),
			})

			// Sign the token using the secret key and return the token string and error if any
			return token.SignedString(secretKey)
		}