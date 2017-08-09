package config

const defaultApplicationPort string = "1234"

// Config ...
type Config interface {
	GetPort() string
	GetDatabaseDSN() string
	GetDebugEnabled() bool
}
