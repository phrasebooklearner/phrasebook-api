package config

import "os"

// NewEnvConfig ...
func NewEnvConfig() *envConfig {
	return &envConfig{}
}

type envConfig struct{}

// GetDatabaseDSN ...
func (ec *envConfig) GetPort() string {
	port := os.Getenv("APPLICATION_PORT")
	if port == "" {
		port = defaultApplicationPort
	}

	return port
}

// GetDatabaseDSN ...
func (ec *envConfig) GetDatabaseDSN() string {
	//return "mysql://dev:dev@tcp(db:3306)/dev"
	return os.Getenv("DATABASE_DSN")
}

func (ec *envConfig) GetDebugEnabled() bool {
	return true // TODO get from environment
}
