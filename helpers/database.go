package helpers

import (
	"database/sql"
	"fmt"

	"git.gigamons.de/Gigamons/Kaoiji/constants"

	// MySQL Driver
	_ "github.com/go-sql-driver/mysql"
)

// Connect to MySQL Database
func Connect(conf constants.Config) (*sql.DB, error) {
	c := conf.MySQL
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%v)/%s", c.Username, c.Password, c.Hostname, c.Port, c.Database))
	return db, err
}

// CheckConnection if database can connect return nil
func CheckConnection(db *sql.DB) error {
	return db.Ping()
}
