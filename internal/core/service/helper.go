package service

import "github.com/tommjj/go-blog-api/internal/logger"

// logIfErr log if error not nil, use logger.Error
func logIfErr(err error) {
	if err != nil {
		logger.Error(err.Error())
	}
}
