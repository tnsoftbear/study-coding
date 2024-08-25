package config

import "os"

// Config stores configuration data for the whole app.
type Config struct {
	// Listening address of the web-server in "host:port" format.
	ListenAddr string
}

// Init reads config from environment.
func Init() Config {
	var config Config
	addr, exists := os.LookupEnv("LISTEN_ADDR")
	if exists {
		config.ListenAddr = addr
	} else {
		config.ListenAddr = "localhost:8080"
	}
	return config
}
