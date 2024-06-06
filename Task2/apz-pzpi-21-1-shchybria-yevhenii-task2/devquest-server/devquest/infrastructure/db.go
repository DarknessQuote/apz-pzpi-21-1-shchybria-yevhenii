package infrastructure

import "database/sql"

type Database interface {
	GetDB() *sql.DB
}