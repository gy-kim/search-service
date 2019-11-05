package main

import (
	"context"

	"github.com/gy-kim/search-service/config"
	"github.com/gy-kim/search-service/internal/auth"
	"github.com/gy-kim/search-service/internal/data"
	"github.com/gy-kim/search-service/internal/list"
	"github.com/gy-kim/search-service/restful"
)

func main() {
	ctx := context.Background()

	server := initializeServer()
	server.Listen(ctx.Done())
}

func initializeServer() *restful.Server {
	lister := list.NewLister(config.App)
	auth := auth.NewJWTAuth(config.App)

	server := restful.New(config.App, lister, auth)
	return server
}

func init() {
	_ = data.NewDAO(config.App)
}
