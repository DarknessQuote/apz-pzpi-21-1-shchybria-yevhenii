package infrastructure

import (
	"database/sql"
	"devquest-server/config"
	"time"
)

type Database interface {
	GetDB() *sql.DB
	GetDBTimeout() time.Duration
	CreateBackup(*config.Config) error
}