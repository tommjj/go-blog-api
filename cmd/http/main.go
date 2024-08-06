package main

import (
	"log"

	"github.com/tommjj/go-blog-api/internal/adapter/http"
	"github.com/tommjj/go-blog-api/internal/config"
	"github.com/tommjj/go-blog-api/internal/logger"
)

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatalf("load config error ::%v", err)
	}

	// setup logger
	err = logger.Set(*config.Logger)
	if err != nil {
		log.Fatalf("set logger err ::%v", err)
	}
	defer logger.Sync()

	r, err := http.New(config.Http)
	if err != nil {
		log.Fatalf("new http err ::%v", err)

	}
	r.Serve()

	logger.Info("setup done!")
}
