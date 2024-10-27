package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresAdapter struct {
	DB *sql.DB
}

func (p *PostgresAdapter) Connect(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	p.DB = db
	return db, nil
}

func (p *PostgresAdapter) Close() error {
	return p.DB.Close()
}

func (p *PostgresAdapter) Insert(query string, args ...interface{}) (sql.Result, error) {
	return p.DB.Exec(query, args...)
}

func (p *PostgresAdapter) Select(query string, args ...interface{}) (*sql.Rows, error) {
	return p.DB.Query(query, args...)
}

func (p *PostgresAdapter) Update(query string, args ...interface{}) (sql.Result, error) {
	return p.DB.Exec(query, args...)
}

func (p *PostgresAdapter) Delete(query string, args ...interface{}) (sql.Result, error) {
	return p.DB.Exec(query, args...)
}

func (p *PostgresAdapter) CreateTable(name string, column string) (sql.Result, error) {
	query := fmt.Sprintf("CREATE TABLE %s(%s)", name, column)
	return p.DB.Exec(query)
}

func (p *PostgresAdapter) CreateDatabase(name string) (sql.Result, error) {
	query := fmt.Sprintf("CREATE DATABASE %s", name)
	return p.DB.Exec(query)
}
