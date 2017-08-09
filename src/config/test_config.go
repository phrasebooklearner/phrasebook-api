package config

type testConfig struct {
	Config
}

func NewTestConfig() Config {
	return &testConfig{
		Config: NewEnvConfig(),
	}
}
