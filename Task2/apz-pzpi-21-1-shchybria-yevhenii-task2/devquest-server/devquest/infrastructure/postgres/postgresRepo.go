package postgres

import (
	"database/sql"
	"devquest-server/config"
	"fmt"
	"sync"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type PostgresRepo struct {
	Db *sql.DB
}

var (
	once       sync.Once
	dbInstance *PostgresRepo
	dbError    error
)

func NewPostgresDB(conf *config.Config) (*PostgresRepo, error) {
	once.Do(func() {
		dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s timezone=%s connect_timeout=%d",
			conf.Database.Host,
			conf.Database.Port,
			conf.Database.User,
			conf.Database.Password,
			conf.Database.DBName,
			conf.Database.SSLMode,
			conf.Database.TimeZone,
			conf.Database.ConnectTimeout,
		)

		db, err := sql.Open("pgx", dsn)
		if err != nil {
			dbError = err
			return
		}

		if err = db.Ping(); err != nil {
			dbError = err
			return
		}

		dbInstance = &PostgresRepo{Db: db}
	})

	return dbInstance, dbError
}

func (p *PostgresRepo) GetDB() *sql.DB {
	return p.Db
}