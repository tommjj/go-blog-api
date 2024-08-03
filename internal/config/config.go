package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		App    *App
		Logger *Logger
		DB     *DB
		Auth   *Auth
	}

	App struct {
		Name string
		Env  string
	}

	Logger struct {
		Level      string
		FileName   string
		Encoder    string
		MaxSize    int
		MaxBackups int
		MaxAge     int
	}

	DB struct {
		FileName string
	}

	Auth struct {
		SecretKey string
		Duration  string
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

	maxSize, err := strconv.Atoi(os.Getenv("LOG_MAX_SIZE"))
	if err != nil {
		return nil, err
	}
	maxBackups, err := strconv.Atoi(os.Getenv("LOG_MAX_BACKUPS"))
	if err != nil {
		return nil, err
	}
	maxAge, err := strconv.Atoi(os.Getenv("LOG_MAX_AGE"))
	if err != nil {
		return nil, err
	}

	logger := &Logger{
		Level:      os.Getenv("LOG_LEVEL"),
		FileName:   os.Getenv("LOG_FILE"),
		Encoder:    os.Getenv("LOG_ENCODER"),
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
	}

	db := &DB{
		FileName: os.Getenv("DB_FILE_NAME"),
	}

	auth := &Auth{
		SecretKey: os.Getenv("AUTH_SECRET"),
		Duration:  os.Getenv("AUTH_TOKEN_DURATION"),
	}

	return &Config{
		App:    app,
		Logger: logger,
		DB:     db,
		Auth:   auth,
	}, nil
}
