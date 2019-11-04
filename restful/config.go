package restful

import "github.com/gy-kim/search-service/logging"

// Config is the configuration for restful package
type Config interface {
	BindServicePort() string
	Logger() logging.Logger
}
