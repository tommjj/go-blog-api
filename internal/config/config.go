package config

import (
	"os"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		App    *App
		Logger *Logger
	}

	App struct {
		Name string
		Env  string
	}

	Logger struct {
		Level string
	}
)

func New() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	app := &App{
		Name: os.Getenv("APP_NAME"),
		Env:  os.Getenv("APP_ENV"),
	}

	logger := &Logger{
		Level: os.Getenv("LOG_LEVEL"),
	}

	return &Config{
		App:    app,
		Logger: logger,
	}, nil
}
