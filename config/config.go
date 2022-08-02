package config

import (
	"os"
)

type Configs struct {
	Database DatabaseConfig
	App      AppConfig
	SMTP     SMTPConfig
}

type DatabaseConfig struct {
	Host     string
	User     string
	Password string
	Database string
	Port     string
}

type AppConfig struct {
	VerificationLength string
	JWTSecret          string
	TokenExpire        string
}

type SMTPConfig struct {
	Email    string
	Password string
	Host     string
	Port     string
}

var Config Configs

func InitConfig() {
	Config.Database = DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
		Port:     os.Getenv("DB_PORT"),
	}

	Config.App = AppConfig{
		VerificationLength: os.Getenv("VERIFICATION_LENGTH"),
		JWTSecret:          os.Getenv("JWT_SECRET"),
		TokenExpire:        os.Getenv("TOKEN_EXPIRE"),
	}

	Config.SMTP = SMTPConfig{
		Email:    os.Getenv("SMTP_EMAIL"),
		Password: os.Getenv("SMTP_PASSWORD"),
		Host:     os.Getenv("SMTP_HOST"),
		Port:     os.Getenv("SMTP_PORT"),
	}
}
