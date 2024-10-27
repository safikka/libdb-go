package db

import "database/sql"

type Database interface {
	Connect(dsn string) (*sql.DB, error)
	Close() error
	Select(query string, args ...interface{}) (*sql.Rows, error)
	Insert(query string, args ...interface{}) (sql.Result, error)
	Update(query string, args ...interface{}) (sql.Result, error)
	Delete(query string, args ...interface{}) (sql.Result, error)
	CreateTable(name string, column string) (sql.Result, error)
	CreateDatabase(query string) (sql.Result, error)
}
