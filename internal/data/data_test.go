package data

import "github.com/gy-kim/search-service/logging"

type testConfig struct {
	url string
}

func (t *testConfig) Logger() logging.Logger {
	return &logging.LoggerStdOut{}
}

func (t *testConfig) DataURL() string {
	return t.url
}
