package main

import (
	"fmt"
	"github.com/cyanidee/bancho-go/handlers"
	"github.com/cyanidee/bancho-go/helpers"
	"log"
	"net/http"
	"os"
)

func init() {
	err, _, created := helpers.ReadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	if created {
		fmt.Println("I've just created a config for you! Please edit settings.toml.")
		os.Exit(0)
	}
}

func main() {
	err, conf, _ := helpers.ReadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	helpers.GlobalConfig = &conf

	if _, err := helpers.ConnectMySQL(conf.MySQL.Hostname, conf.MySQL.Port,
									  conf.MySQL.Username, conf.MySQL.Password,
									  conf.MySQL.Database); err != nil {
		log.Fatalln(err)
	}

	defer helpers.DBConn.Close()

	if err = helpers.DBConn.Ping(); err != nil {
		log.Fatalln(err)
	}

	http.HandleFunc("/", handlers.Handle)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", conf.Server.Hostname, conf.Server.Port), nil))
}
