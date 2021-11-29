package bdd

import "fmt"

type DatabaseCredentials struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         int
}

func NewDatabaseCredentials(host string, username string, password string, databaseName string, port int) *DatabaseCredentials {
	databaseCredentials := new(DatabaseCredentials)
	databaseCredentials.Host = host
	databaseCredentials.Username = username
	databaseCredentials.Password = password
	databaseCredentials.DatabaseName = databaseName
	databaseCredentials.Port = port
	return databaseCredentials
}

func (d DatabaseCredentials) String() string {
	return fmt.Sprintf(d.Username+":"+d.Password+"@tcp("+d.Host+":%d)/"+d.DatabaseName, d.Port)
}
