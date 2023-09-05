package providers

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnectDB function to create a connection to database
func ConnectDB(driver string, dsn string) (*gorm.DB, error) {
	// Open a connection to database using the driver and dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// Return the db object and error if any
	return db, nil
}
