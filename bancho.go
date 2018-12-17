package main

import (
	"fmt"
	"github.com/Gigamons/Kaoiji/handlers"
	"github.com/Gigamons/Kaoiji/helpers"
	"github.com/valyala/fasthttp"
	"log"
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


func RunHTTPServer(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/":
		handlers.HandleRoot(ctx)
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

	fmt.Println(fmt.Sprintf("Server should be listening at port %d", conf.Server.Port))
	log.Fatalln(fasthttp.ListenAndServe(fmt.Sprintf("%s:%d", conf.Server.Hostname, conf.Server.Port), RunHTTPServer))
}
