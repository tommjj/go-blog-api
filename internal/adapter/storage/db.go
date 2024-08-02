package storage

import (
	"database/sql"

	"github.com/google/uuid"
	sqliteGo "github.com/mattn/go-sqlite3"
	"github.com/tommjj/go-blog-api/internal/adapter/storage/schema"
	"github.com/tommjj/go-blog-api/internal/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const CustomDriverName = "sqlite3_extended"

type DB struct {
	*gorm.DB
}

func setSqliteCustomDriver() {
	sql.Register(CustomDriverName,
		&sqliteGo.SQLiteDriver{
			ConnectHook: func(conn *sqliteGo.SQLiteConn) error {
				err := conn.RegisterFunc(
					"gen_random_uuid",
					func(arguments ...interface{}) (string, error) {
						return uuid.NewString(), nil // Return a string value.
					},
					true,
				)
				return err
			},
		},
	)
}

func New(conf config.DB) (*DB, error) {
	setSqliteCustomDriver()

	conn, err := sql.Open(CustomDriverName, conf.FileName)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(sqlite.Dialector{
		DriverName: CustomDriverName,
		DSN:        conf.FileName,
		Conn:       conn,
	}, &gorm.Config{
		SkipDefaultTransaction:   true,
		DisableNestedTransaction: true,
	})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&schema.User{}, &schema.Blog{})
	if err != nil {
		return nil, err
	}

	return &DB{
		db,
	}, nil
}
