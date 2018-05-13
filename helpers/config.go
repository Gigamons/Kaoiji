package helpers

import (
	"io/ioutil"

	"git.gigamons.de/Gigamons/Kaoiji/constants"
	"gopkg.in/yaml.v2"
)

func CreateConfig() {
	conf := constants.Config{}
	conf.MySQL.Hostname = "127.0.0.1"
	conf.MySQL.Port = 3306
	conf.MySQL.Database = "gigamons"
	conf.Server.Hostname = "127.0.0.1"
	conf.Server.Port = 5001
	conf.Server.Debug = false
	conf.Redis.Hostname = "127.0.0.1"
	conf.Redis.Port = 5001

	c, err := yaml.Marshal(&conf)
	if err != nil {
		panic(err)
	}
	if ioutil.WriteFile("config.yml", c, 0644) != nil {
		panic(err)
	}
}

func GetConfig() constants.Config {
	f, err := ioutil.ReadFile("config.yml")
	if err != nil {
		panic(err)
	}
	c := constants.Config{}
	err = yaml.Unmarshal(f, &c)
	if err != nil {
		panic(err)
	}
	return c
}
