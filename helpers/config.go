package helpers

import (
	"github.com/BurntSushi/toml"
	"math/rand"
	"os"
	"strings"
	"time"
)

type DatabaseConfig struct {
	Hostname string
	Port uint16

	Username string
	Password string

	Database string
}

type ServerConfig struct {
	Hostname string
	Port uint16
}

type SecurityConfig struct {
	Secret string
	SecretSeed int
	Syntax string
}

type Config struct {
	Server ServerConfig
	Security SecurityConfig
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

var asciichars = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!\"§$%&/()=?`´#+*~'")

func randString() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 64)
	for i := range b {
		b[i] = asciichars[rand.Intn(len(asciichars))]
	}
	return string(b)
}

func ReadConfig() (err error, conf Config, created bool) {
	if _, err := toml.DecodeFile("settings.toml", &conf); err != nil {
		conf.MySQL.Hostname = "localhost"
		conf.MySQL.Port = 3306
		conf.MySQL.Username = "root"

		conf.Server.Hostname = "localhost"
		conf.Server.Port = 45471

		conf.Security.SecretSeed = rand.Int()
		time.Sleep(5)
		conf.Security.Secret = randString()
		time.Sleep(5)
		conf.Security.Syntax = randString()
		conf.Security.Syntax = strings.Replace(conf.Security.Syntax, "s", "%secret%", -1)
		conf.Security.Syntax = strings.Replace(conf.Security.Syntax, "D", "%secret%", -1)
		conf.Security.Syntax = strings.Replace(conf.Security.Syntax, "f", "%secret%", -1)
		conf.Security.Syntax = strings.Replace(conf.Security.Syntax, "c", "%secret%", -1)
		conf.Security.Syntax = strings.Replace(conf.Security.Syntax, "p", "%password%", -1)
		conf.Security.Syntax = strings.Replace(conf.Security.Syntax, "G", "%password%", -1)
		conf.Security.Syntax = strings.Replace(conf.Security.Syntax, "h", "%password%", -1)
		conf.Security.Syntax = strings.Replace(conf.Security.Syntax, "#", "%password%", -1)

		if !strings.Contains(conf.Security.Syntax, "%secret%") {
			conf.Security.Syntax += "%secret%"
		}

		if !strings.Contains(conf.Security.Syntax, "%password%") {
			conf.Security.Syntax += "%password%"
		}

		err = WriteConfig(&conf)
		created = true
	}

	return
}