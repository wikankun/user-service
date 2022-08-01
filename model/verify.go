package model

import (
	"github.com/wikankun/user-service/database"
)

type Verification struct {
	ID     int    `gorm:"not null" json:"id"`
	UserID int    `gorm:"not null" json:"user_id"`
	Code   string `gorm:"type:varchar(100);not null" json:"code"`
	User   User
}

func CreateVerification(id int, code string) error {
	verification := Verification{
		UserID: id,
		Code:   code,
	}

	result := database.Connector.Create(&verification)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func VerifyCode(code string) (Verification, error) {
	var verification Verification

	result := database.Connector.Preload("User").Where("code = ?", code).First(&verification)
	if result.Error != nil {
		return verification, result.Error
	}

	return verification, nil
}

func DeleteVerifyCode(code string) error {
	var verification Verification

	result := database.Connector.Where("code = ?", code).Delete(&verification)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
