package main

import (
	"os"

	"github.com/Gigamons/common/consts"
	"github.com/Gigamons/common/logger"

	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Gigamons/Kaoiji/global"
	"github.com/Gigamons/Kaoiji/server"
	"github.com/Gigamons/common/helpers"
)

func init() {
	if _, err := os.Stat("config.yml"); os.IsNotExist(err) {
		helpers.CreateConfig("config", constants.Config{MySQL: consts.MySQLConf{Database: "gigamons", Hostname: "localhost", Port: 3306, Username: "root"}})
		logger.Infoln("I've just created a config.yml! please edit!")
		os.Exit(0)
	}
}

func main() {
	var err error
	var conf constants.Config

	helpers.GetConfig("config", &conf)

	helpers.Connect(&conf.MySQL)
	if err = helpers.DB.Ping(); err != nil {
		panic(err)
	}
	global.CONFIG = &conf

	if conf.Server.Debug {
		os.Setenv("DEBUG", "true")
	}

	defer helpers.DB.Close()

	server.StartServer(conf.Server.Port)
}
