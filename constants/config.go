package constants

// Config is for the config.yml file
type Config struct {
	Server struct {
		Hostname   string
		Port       int
		FreeDirect bool
		Debug      bool
	}
	MySQL struct {
		Hostname string
		Port     int
		Username string
		Password string
		Database string
	}
	Redis struct {
		Hostname string
		Port     int
	}
}
