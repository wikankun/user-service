package database

import (
	"fmt"
	"log"

	"github.com/wikankun/user-service/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var GetConnectionString = func() string {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.Config.Database.Host,
		config.Config.Database.User,
		config.Config.Database.Password,
		config.Config.Database.Database,
		config.Config.Database.Port,
	)

	return connectionString
}

// Connector variable used for CRUD operation's
var Connector *gorm.DB

// Connect creates MySQL connection
func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return err
	}
	log.Println("Connected to Database")
	return nil
}
