package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		App    *App
		Logger *Logger
		DB     *DB
		Auth   *Auth
		Http   *Http
		Redis  *Redis
	}

	App struct {
		Name string
		Env  string
	}

	Logger struct {
		Level         string
		Encoder       string
		LogFileWriter *LogFileWriter
	}

	LogFileWriter struct {
		FileName   string
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

	Http struct {
		Env            string
		AllowedOrigins []string
		URL            string
		Port           int
		Logger         Logger
	}

	Redis struct {
		Addr     string
		Password string
	}
)

func New() (*Config, error) {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			return nil, err
		}
	}

	app := GetAppConf()

	logger, err := GetLoggerConf()
	if err != nil {
		return nil, err
	}

	db := GetDBConf()

	auth := GetAuthConf()

	http, err := GetHTTPConf()
	if err != nil {
		return nil, err
	}

	redis := GetRedisConf()

	return &Config{
		App:    app,
		Logger: logger,
		DB:     db,
		Auth:   auth,
		Http:   http,
		Redis:  redis,
	}, nil
}

func GetAppConf() *App {
	return &App{
		Name: os.Getenv("APP_NAME"),
		Env:  os.Getenv("APP_ENV"),
	}
}

func GetLoggerConf() (*Logger, error) {
	if os.Getenv("LOG_ENABLE_FILE_WRITER") != "true" {
		return &Logger{
			Level:         os.Getenv("LOG_LEVEL"),
			Encoder:       os.Getenv("LOG_ENCODER"),
			LogFileWriter: nil,
		}, nil
	}

	maxSize, err := strconv.Atoi(os.Getenv("LOG_MAX_SIZE"))
	if err != nil {
		return nil, fmt.Errorf("LOG_MAX_SIZE must to be a number: %v", err)
	}
	maxBackups, err := strconv.Atoi(os.Getenv("LOG_MAX_BACKUPS"))
	if err != nil {
		return nil, fmt.Errorf("LOG_MAX_BACKUPS must to be a number: %v", err)
	}
	maxAge, err := strconv.Atoi(os.Getenv("LOG_MAX_AGE"))
	if err != nil {
		return nil, fmt.Errorf("LOG_MAX_AGE must to be a number: %v", err)
	}

	return &Logger{
		Level:   os.Getenv("LOG_LEVEL"),
		Encoder: os.Getenv("LOG_ENCODER"),
		LogFileWriter: &LogFileWriter{
			FileName:   os.Getenv("LOG_FILE"),
			MaxSize:    maxSize,
			MaxBackups: maxBackups,
			MaxAge:     maxAge,
		},
	}, nil
}

func GetDBConf() *DB {
	return &DB{
		FileName: os.Getenv("DB_FILE_NAME"),
	}
}

func GetAuthConf() *Auth {
	return &Auth{
		SecretKey: os.Getenv("AUTH_SECRET"),
		Duration:  os.Getenv("AUTH_TOKEN_DURATION"),
	}
}

func GetHTTPConf() (*Http, error) {
	allowedOrigins := strings.Split(os.Getenv("HTTP_ALLOWED_ORIGINS"), ",")
	port, err := strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		return nil, fmt.Errorf("HTTP_PORT must to be a number: %v", err)
	}

	logger := Logger{
		Level:         os.Getenv("HTTP_LOG_LEVEL"),
		Encoder:       os.Getenv("HTTP_LOG_ENCODER"),
		LogFileWriter: nil,
	}

	if os.Getenv("HTTP_LOG_ENABLE_FILE_WRITER") == "true" {
		maxSize, err := strconv.Atoi(os.Getenv("HTTP_LOG_MAX_SIZE"))
		if err != nil {
			return nil, fmt.Errorf("HTTP_LOG_MAX_SIZE must to be a number: %v", err)
		}
		maxBackups, err := strconv.Atoi(os.Getenv("HTTP_LOG_MAX_BACKUPS"))
		if err != nil {
			return nil, fmt.Errorf("HTTP_LOG_MAX_BACKUPS must to be a number: %v", err)
		}
		maxAge, err := strconv.Atoi(os.Getenv("HTTP_LOG_MAX_AGE"))
		if err != nil {
			return nil, fmt.Errorf("HTTP_LOG_MAX_AGE must to be a number: %v", err)
		}

		logger.LogFileWriter = &LogFileWriter{
			FileName:   os.Getenv("HTTP_LOG_FILE"),
			MaxSize:    maxSize,
			MaxBackups: maxBackups,
			MaxAge:     maxAge,
		}
	}

	return &Http{
		Env:            os.Getenv("APP_ENV"),
		AllowedOrigins: allowedOrigins,
		URL:            os.Getenv("HTTP_URL"),
		Port:           port,
		Logger:         logger,
	}, nil
}

func GetRedisConf() *Redis {
	return &Redis{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASS"),
	}
}
