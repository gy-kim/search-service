package config

import (
	"fmt"
	"os"

	"github.com/gy-kim/search-service/logging"
)

const (
	envDataURL     = "ELASTIC_URL"
	envServicePort = "SERVICE_PORT"
)

const (
	defDataURL     = "http://127.0.0.1:9200"
	defServicePort = "9000"
)

// App is the application config
var App *Config

// Config defines the configuration of app
type Config struct {
	servicePort string

	dataURL string

	logger logging.Logger
}

// BindServicePort returns the host and port this service.
func (c *Config) BindServicePort() string {
	return fmt.Sprintf(":%s", c.servicePort)
}

// Logger returns a reference of logger
func (c *Config) Logger() logging.Logger {
	if c.logger == nil {
		c.logger = &logging.LoggerStdOut{}
	}
	return c.logger
}

// DataURL returns dataURL for elastic.
func (c *Config) DataURL() string {
	return c.dataURL
}

func init() {
	App = &Config{}

	App.servicePort = getEnvWithDefault(envServicePort, defServicePort)
	App.dataURL = getEnvWithDefault(envDataURL, defDataURL)
}

func getEnvWithDefault(key string, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}
