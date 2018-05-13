package main

import (
	"fmt"
	"os"

	"git.gigamons.de/Gigamons/Kaoiji/global"
	"git.gigamons.de/Gigamons/Kaoiji/helpers"
	"git.gigamons.de/Gigamons/Kaoiji/server"
)

func init() {
	if _, err := os.Stat("config.yml"); os.IsNotExist(err) {
		helpers.CreateConfig()
		fmt.Println("I've just created a config.yml! please edit!")
		os.Exit(0)
	}
}

func main() {
	var err error

	global.DB, err = helpers.Connect(helpers.GetConfig())
	if err != nil {
		panic(err)
	}
	if err = helpers.CheckConnection(global.DB); err != nil {
		panic(err)
	}
	global.CONFIG = helpers.GetConfig()

	defer global.DB.Close()

	server.StartServer(5001)
}
