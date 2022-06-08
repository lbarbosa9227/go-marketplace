package database

import (
	"crudapi/Internal/logs"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type MySqlClient struct {
	*sql.DB
}

func NewSqlClient(source string) *MySqlClient {

	db, err := sql.Open("mysql", source)
	if err != nil {
		logs.Log().Errorf("canot create db tentat: %s", err.Error())
		panic("error in database connection")
	}

	err = db.Ping()

	if err != nil {
		logs.Log().Warn("canot connect  to database")

	}
	return &MySqlClient{db}

}

func (c *MySqlClient) ViewStats() sql.DBStats {
	return c.Stats()

}
