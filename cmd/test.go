package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/tommjj/go-blog-api/internal/adapter/storage/sqlite"
	"github.com/tommjj/go-blog-api/internal/adapter/storage/sqlite/repository"
	"github.com/tommjj/go-blog-api/internal/config"
	"github.com/tommjj/go-blog-api/internal/core/ports"
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

	db, err := sqlite.New(*config.DB)
	fatalIfErr(err)

	var repo ports.IBlogRepository = repository.NewBlogRepository(db)

	blog, err := repo.GetBlogsByAuthorID(context.TODO(), uuid.MustParse("539f68d9-438b-4e96-a27f-17e9aabd152c"), 1, 10)
	fatalIfErr(err)

	fmt.Println(blog)
}
