package helpers

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DBConn *sql.DB

func AntiTimeout() {
	for {
		if DBConn == nil {
			break
		} else if err := DBConn.Ping(); err != nil {
			log.Fatalln(err)
			break
		}

		time.Sleep(time.Second * 30)
	}
}

func ConnectMySQL(hostname string, port uint16,
	         username string, password string,
	         database string) (db *sql.DB, err error) {
	conStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, hostname, port, database)
	db, err = sql.Open("mysql", conStr)

	DBConn = db
	go AntiTimeout()
	return // returns db and err.
}