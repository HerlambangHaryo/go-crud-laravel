package main

import (
  "fmt"

  "github.com/gofiber/fiber/v2"
  "github.com/HerlambangHaryo/go-crud-laravel/app/providers"
  "github.com/HerlambangHaryo/go-crud-laravel/config"
  "github.com/HerlambangHaryo/go-crud-laravel/database/migrations"
  "github.com/HerlambangHaryo/go-crud-laravel/routes"
  "github.com/gofiber/template/html/v2"
)

// Global variable for database connection
var db *gorm.DB

func main() {
  // Load environment variables from .env file
  config.LoadEnv()

  // Get database driver and dsn from environment variables
  driver := config.GetEnv("DB_DRIVER")
  dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
  	config.GetEnv("DB_USER"),
  	config.GetEnv("DB_PASSWORD"),
  	config.GetEnv("DB_HOST"),
  	config.GetEnv("DB_PORT"),
  	config.GetEnv("DB_NAME"),
  )

  // Connect to database using driver and dsn
  var err error
  db, err = providers.ConnectDB(driver, dsn)
  if err != nil {
  	panic(err)
  }

  // Create users table in database
  err = migrations.CreateUsersTable(db)
  if err != nil {
  	panic(err)
  }

  // Create a fiber app
  app := fiber.New()

  // Create a new html engine
	engine := html.New("./resources/views", ".gohtml")

	// Register the html engine to the app
	app.Views(engine)

  // Register web routes
  routes.WebRoutes(app)

  // Register api routes
  routes.ApiRoutes(app)

}
