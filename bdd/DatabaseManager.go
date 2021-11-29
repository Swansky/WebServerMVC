package bdd

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type DatabaseManager struct {
	credentials *DatabaseCredentials
}

func NewDatabaseManager(credentials *DatabaseCredentials) *DatabaseManager {
	databaseManager := new(DatabaseManager)
	databaseManager.credentials = credentials
	return databaseManager
}

func (d *DatabaseManager) CreateConnection() *sql.DB {
	db, err := sql.Open("mysql", d.credentials.String())
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
