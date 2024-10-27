package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteAdapter struct {
	DB *sql.DB
}

func (s *SQLiteAdapter) Connect(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}
	s.DB = db
	return db, nil
}

func (s *SQLiteAdapter) Close() error {
	return s.DB.Close()
}

func (s *SQLiteAdapter) Insert(query string, args ...interface{}) (sql.Result, error) {
	return s.DB.Exec(query, args...)
}

func (s *SQLiteAdapter) Select(query string, args ...interface{}) (*sql.Rows, error) {
	return s.DB.Query(query, args...)
}

func (s *SQLiteAdapter) Update(query string, args ...interface{}) (sql.Result, error) {
	return s.DB.Exec(query, args...)
}

func (s *SQLiteAdapter) Delete(query string, args ...interface{}) (sql.Result, error) {
	return s.DB.Exec(query, args...)
}

func (s *SQLiteAdapter) CreateTable(name string, column string) (sql.Result, error) {
	query := fmt.Sprintf("CREATE TABLE %s(%s)", name, column)
	return s.DB.Exec(query)
}

func (s *SQLiteAdapter) CreateDatabase(name string) (sql.Result, error) {
	return nil, fmt.Errorf("SQLite does not support dynamic database creation")
}
