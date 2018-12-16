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

func WriteConfig(conf *Config) (err error) {
	f, err := os.Create("settings.toml")
	if err != nil {
		return
	}
	defer f.Close()

	return toml.NewEncoder(f).Encode(conf)
}

func ReadConfig() (err error, conf Config) {
	if _, err := toml.DecodeFile("settings.toml", &conf); err != nil {
		conf.MySQL.Hostname = "localhost"
		conf.MySQL.Port = 3306
		conf.MySQL.Username = "root"

		err = WriteConfig(&conf)
	}

	return
}