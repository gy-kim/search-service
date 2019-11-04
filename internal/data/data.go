package data

import (
	"github.com/gy-kim/search-service/logging"
	"github.com/olivere/elastic"
)

var (
	client *elastic.Client
)

// Config is the configuration for data apckage
type Config interface {
	DataURL() string
	Logger() logging.Logger
}

var getClient = func(cfg Config) (*elastic.Client, error) {
	if client == nil {
		var err error
		client, err = elastic.NewClient(
			elastic.SetURL(cfg.DataURL()),
			elastic.SetSniff(false),
		)
		if err != nil {
			cfg.Logger().Error("failed to connect elastic server. err: %v", err)
			panic(err)
		}
		cfg.Logger().Info("success to connect elastic server.")
	}
	return client, nil
}
