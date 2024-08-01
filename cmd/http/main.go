package main

import (
	"log"

	"github.com/tommjj/go-blog-api/internal/config"
	"github.com/tommjj/go-blog-api/internal/logger"
)

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatalf("load config error ::%v", err)
	}

	err = logger.Set(*config.Logger)
	if err != nil {
		log.Fatalf("set logger err ::%v", err)
	}

	logger.L.Info("setup done!")
}
