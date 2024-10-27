// file: db/mysql_adapter.go
package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLAdapter struct {
	DB *sql.DB
}

func (m *MySQLAdapter) Connect(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	m.DB = db
	return db, nil
}

func (m *MySQLAdapter) Close() error {
	return m.DB.Close()
}

func (m *MySQLAdapter) Insert(query string, args ...interface{}) (sql.Result, error) {
	return m.DB.Exec(query, args...)
}

func (m *MySQLAdapter) Select(query string, args ...interface{}) (*sql.Rows, error) {
	return m.DB.Query(query, args...)
}

func (m *MySQLAdapter) Update(query string, args ...interface{}) (sql.Result, error) {
	return m.DB.Exec(query, args...)
}

func (m *MySQLAdapter) Delete(query string, args ...interface{}) (sql.Result, error) {
	return m.DB.Exec(query, args...)
}

func (m *MySQLAdapter) CreateTable(name string, column string) (sql.Result, error) {
	query := fmt.Sprintf("CREATE TABLE %s(%s)", name, column)
	return m.DB.Exec(query)
}

func (m *MySQLAdapter) CreateDatabase(name string) (sql.Result, error) {
	query := fmt.Sprintf("CREATE DATABASE %s", name)
	return m.DB.Exec(query)
}
