package constants

import (
	"github.com/Gigamons/common/consts"
)

// Config is for the config.yml file
type Config struct {
	Server struct {
		Hostname   string
		Port       int
		FreeDirect bool
		Debug      bool
	}
	API struct {
		KokoroHost   string
		KokoroAPIKey string
		APIKey       string
	}
	MySQL consts.MySQLConf
	Redis struct {
		Hostname string
		Port     int
	}
}
