package config

import "fmt"

// Config is a placeholder for future configuration.
type Config struct{}

// ConfigLoader loads and parses configuration from a given path.
type ConfigLoader interface {
	Load(path string) (Config, error)
}

// StubConfigLoader is a stub that always returns not implemented.
type StubConfigLoader struct{}

func (s *StubConfigLoader) Load(_ string) (Config, error) {
	return Config{}, fmt.Errorf("not implemented")
}
