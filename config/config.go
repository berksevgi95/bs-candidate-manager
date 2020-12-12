package config

type Config struct {
	Host string
	Port string
	Database string
}

func GetConfig() *Config {
	return &Config{
		Host:  "localhost",
		Port: "27017",
		Database: "Otsimo",
	}
}
