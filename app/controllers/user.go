package controllers

import ( 
	"strconv"
	"github.com/gofiber/fiber/v2"
	"github.com/HerlambangHaryo/go-crud-laravel/app/models"
	"github.com/HerlambangHaryo/go-crud-laravel/app/services"
)

	// Create a new user instance and assign the email
		user := new(models.User)
		user.Email = email

	// Read the user by email from the database
		err := user.ReadByEmail(db)
		if err != nil {
			// If the user is not found, redirect to register page
			return c.Redirect("/register")
		}

	// Check the password with the hashed password from the database
		match := services.CheckPassword(password, user.Password)
		if !match {
			// If the password does not match, redirect to login page
			return c.Redirect("/login")
		}

	// Generate a jwt token for the user using the user id
		token, err := services.GenerateToken(user)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}

	// Set a cookie with the token and redirect to user index page
		c.Cookie(&fiber.Cookie{
			Name:     "token",
			Value:    token,
			Expires:  time.Now().Add(time.Hour * 24),
		})
		
return c.Redirect("/user")

}