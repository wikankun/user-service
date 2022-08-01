package model

import (
	"time"

	"github.com/wikankun/user-service/database"
)

type User struct {
	ID        int       `gorm:"not null" json:"id"`
	Username  string    `gorm:"type:varchar(100);not null" json:"username"`
	Email     string    `gorm:"type:varchar(100);not null" json:"email"`
	Password  string    `gorm:"type:varchar(100);not null" json:"password"`
	CreatedAt time.Time `gorm:"not null" json:"created_at,omitempty"`
	IsAdmin   bool      `gorm:"not null" json:"is_admin,omitempty"`
	Verified  bool      `gorm:"not null" json:"verified,omitempty"`
}

func GetUserWithID(id int) (User, error) {
	user := User{
		ID: id,
	}

	result := database.Connector.First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func GetUserWithUsername(username string) (User, error) {
	user := User{}

	result := database.Connector.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func GetAllUsers() ([]User, error) {
	var users []User

	result := database.Connector.Find(&users)
	if result.Error != nil {
		return users, result.Error
	}
	return users, nil
}

func IsUserAdmin(id int) (bool, error) {
	user := User{
		ID: id,
	}

	result := database.Connector.First(&user)
	if result.Error != nil {
		return false, result.Error
	}
	return user.IsAdmin, nil
}

func AddUser(user User) (int, error) {
	result := database.Connector.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	return user.ID, nil
}

func UpdateUser(user User) error {
	result := database.Connector.Save(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteUser(id int) error {
	user := User{
		ID: id,
	}

	result := database.Connector.Delete(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
