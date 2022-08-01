package migration

import (
	"github.com/wikankun/user-service/database"
	"github.com/wikankun/user-service/model"
)

func MigrateDB() error {
	err := database.Connector.AutoMigrate(
		&model.User{},
		&model.Verification{},
	)
	if err != nil {
		return err
	}
	return nil
}
