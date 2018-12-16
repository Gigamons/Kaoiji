package helpers

import (
	"github.com/BurntSushi/toml"
	"os"
)

type DatabaseConfig struct {
	Hostname string
	Port uint16

	Username string
	Password string

	Database string
}

type Config struct {
	MySQL DatabaseConfig
}

var GlobalConfig *Config

func WriteConfig(conf *Config) (err error) {
	f, err := os.Create("settings.toml")
	if err != nil {
		return
	}
	defer f.Close()

	encoder := toml.NewEncoder(f)
	encoder.Indent = ""
	return encoder.Encode(conf)
}

func ReadConfig() (err error, conf Config, created bool) {
	if _, err := toml.DecodeFile("settings.toml", &conf); err != nil {
		conf.MySQL.Hostname = "localhost"
		conf.MySQL.Port = 3306
		conf.MySQL.Username = "root"

		err = WriteConfig(&conf)
		created = true
	}

	return
}