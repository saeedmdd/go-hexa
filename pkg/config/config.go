package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var AppConfig Config

type Config struct {
	Database Database
	App      App
}

type App struct {
	AppName  string
	AppPort  string
	AppDebug bool
}

type Database struct {
	Postgres Postgres
}

type Postgres struct {
	Host     string
	UserName string
	Password string
	Database string
	Port     string
	Sslmode  string
}

func Apply() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	appDebug, err := strconv.ParseBool(os.Getenv("APP_DEBUG"))
	if err != nil {
		return err
	}

	app := App{
		AppName:  os.Getenv("APP_NAME"),
		AppPort:  os.Getenv("APP_PORT"),
		AppDebug: appDebug,
	}

	postgres := Postgres{
		Host:     os.Getenv("POSTGRES_HOST"),
		UserName: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: os.Getenv("POSTGRES_DATABASE"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Sslmode: os.Getenv("POSTGRES_SSL_MODE"),
	}
	database := Database{
		Postgres: postgres,
	}

	AppConfig.App = app
	AppConfig.Database = database

	return nil
}
