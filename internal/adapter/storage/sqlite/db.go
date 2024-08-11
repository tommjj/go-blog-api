package sqlite

import (
	"database/sql"
	"sync"

	"github.com/google/uuid"
	sqliteGo "github.com/mattn/go-sqlite3"
	"github.com/tommjj/go-blog-api/internal/adapter/storage/sqlite/schema"
	"github.com/tommjj/go-blog-api/internal/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const customDriverName = "sqlite3_extended"

var one = sync.Once{}

type DB struct {
	*gorm.DB
}

func setSqliteCustomDriver() {
	sql.Register(customDriverName,
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
	one.Do(
		setSqliteCustomDriver,
	)

	conn, err := sql.Open(customDriverName, conf.FileName)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(sqlite.Dialector{
		DriverName: customDriverName,
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
