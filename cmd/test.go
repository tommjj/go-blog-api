package main

import (
	"context"

	"github.com/google/uuid"
	"github.com/tommjj/go-blog-api/internal/adapter/storage"
	"github.com/tommjj/go-blog-api/internal/adapter/storage/repository"
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

	db, err := storage.New(*config.DB)
	fatalIfErr(err)

	rp := repository.NewBlogRepository(db)

	err = rp.DeleteBlog(context.TODO(), uuid.MustParse("f68d7e4f-ce94-4346-adec-9b8c7363f93a"))
	fatalIfErr(err)
}
