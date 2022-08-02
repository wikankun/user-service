package config

import (
	"os"
)

type Configs struct {
	App      AppConfig
	Database DatabaseConfig
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
	Port               string
	JWTSecret          string
	TokenExpire        string
	VerificationLength string
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
		Port:               os.Getenv("PORT"),
		JWTSecret:          os.Getenv("JWT_SECRET"),
		TokenExpire:        os.Getenv("TOKEN_EXPIRE"),
		VerificationLength: os.Getenv("VERIFICATION_LENGTH"),
	}

	Config.SMTP = SMTPConfig{
		Email:    os.Getenv("SMTP_EMAIL"),
		Password: os.Getenv("SMTP_PASSWORD"),
		Host:     os.Getenv("SMTP_HOST"),
		Port:     os.Getenv("SMTP_PORT"),
	}
}
