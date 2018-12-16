package helpers

import (
	"fmt"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DBConn *sql.DB

func AntiTimeout() {

}

func ConnectMySQL(hostname string, port uint16,
	         username string, password string,
	         database string) (db *sql.DB, err error) {
	conStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, hostname, port, database)
	db, err = sql.Open("mysql", conStr)

	DBConn = db
	go AntiTimeout()
	return
}