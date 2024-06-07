package infrastructure

import (
	"database/sql"
	"time"
)

type Database interface {
	GetDB() *sql.DB
	GetDBTimeout() time.Duration
}