package db

import "fmt"

func NewDatabaseAdapter(dbType, dsn string) (Database, error) {
	switch dbType {
	case "mysql":
		adapter := &MySQLAdapter{}
		_, err := adapter.Connect(dsn)
		return adapter, err
	case "postgres":
		adapter := &PostgresAdapter{}
		_, err := adapter.Connect(dsn)
		return adapter, err
	case "sqlite":
		adapter := &SQLiteAdapter{}
		_, err := adapter.Connect(dsn)
		return adapter, err
	default:
		return nil, fmt.Errorf("unsupported database type: %s", dbType)
	}
}
