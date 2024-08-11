package main

import (
	"github.com/tommjj/go-blog-api/internal/adapter/http"
	"github.com/tommjj/go-blog-api/internal/config"
	"github.com/tommjj/go-blog-api/internal/logger"
)

func fatalIfErr(err error) {
	if err != nil {
		logger.Fatal(err.Error())
	}
}

func main() {
	config, err := config.New()
	fatalIfErr(err)

	// setup logger
	err = logger.Set(*config.Logger)
	fatalIfErr(err)
	defer logger.Sync()

	r, err := http.New(config.Http)
	fatalIfErr(err)

	r.Serve()
	logger.Info("setup done!")
}
