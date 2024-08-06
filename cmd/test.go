package main

import (
	"context"

	"github.com/google/uuid"
	"github.com/tommjj/go-blog-api/internal/adapter/storage"
	"github.com/tommjj/go-blog-api/internal/adapter/storage/repository"
	"github.com/tommjj/go-blog-api/internal/config"
	"github.com/tommjj/go-blog-api/internal/core/util"
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

	defer logger.L.Sync()

	db, err := storage.New(*config.DB)
	fatalIfErr(err)

	userSv := repository.NewUserRepository(db)

	hashed, err := util.HashPassword("password")
	fatalIfErr(err)

	_, err = userSv.UpdateUserByMap(context.TODO(),
		uuid.MustParse("639f68d9-438b-4e96-a27f-17e9aabd152c"),
		&map[string]interface{}{
			"name":     "Mostima",
			"password": hashed,
		})
	fatalIfErr(err)
}
