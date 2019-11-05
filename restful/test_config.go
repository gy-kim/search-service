package restful

import (
	"fmt"

	"github.com/gy-kim/search-service/logging"
)

type testConfig struct{}

func (t *testConfig) BindServicePort() string {
	return ":9000"
}
func (t *testConfig) Logger() logging.Logger {
	fmt.Println("Logger")
	return &logging.LoggerStdOut{}
}
