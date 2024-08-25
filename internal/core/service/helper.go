package service

import "github.com/tommjj/go-blog-api/internal/logger"

// logOnError log if error not nil, use logger.Error
func logOnError(err error) {
	if err != nil {
		logger.Error(err.Error())
	}
}
