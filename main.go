package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/reversTeam/go-ms-skeleton/services/ping"
	"github.com/reversTeam/go-ms-tools/middlewares"
	"github.com/reversTeam/go-ms-tools/services/abs"
	"github.com/reversTeam/go-ms-tools/services/child"
	"github.com/reversTeam/go-ms-tools/services/people"
	"github.com/reversTeam/go-ms-tools/services/signin"
	"github.com/reversTeam/go-ms/core"
)

const (
	GO_MS_CONFIG_FILEPATH = "./config/config.yml"
)

var (
	configFilePath = flag.String("config", GO_MS_CONFIG_FILEPATH, "yaml config filepath")
)

type TestMiddleware struct {
	core.BaseMiddleware
}

func (t *TestMiddleware) Apply(ctx context.Context, req interface{}) (context.Context, interface{}, error) {
	fmt.Println("TEST MIDDLEWARE IS APPLIED")

	return ctx, req, nil
}

var goMsServices = map[string]core.GoMsServiceFunc{
	"abs": core.RegisterServiceMap(func(ctx *core.Context, name string, config core.ServiceConfig) core.GoMsServiceInterface {
		return abs.NewService(ctx, name, config)
	}),
	"child": core.RegisterServiceMap(func(ctx *core.Context, name string, config core.ServiceConfig) core.GoMsServiceInterface {
		return child.NewService(ctx, name, config)
	}),
	"ping": core.RegisterServiceMap(func(ctx *core.Context, name string, config core.ServiceConfig) core.GoMsServiceInterface {
		return ping.NewService(ctx, name, config)
	}),
	"people": core.RegisterServiceMap(func(ctx *core.Context, name string, config core.ServiceConfig) core.GoMsServiceInterface {
		return people.NewService(ctx, name, config)
	}),
	"signin": core.RegisterServiceMap(func(ctx *core.Context, name string, config core.ServiceConfig) core.GoMsServiceInterface {
		return signin.NewService(ctx, name, config)
	}),
}
var goMsMiddleWare = map[string]core.Middleware{
	"AuthMiddleware":   &middlewares.AuthMiddleware{},
	"UnAuthMiddleware": &middlewares.UnauthenticatedMiddleware{},
	"TestMiddleware":   &TestMiddleware{},
}

func main() {
	flag.Parse()
	config, err := core.NewConfig(*configFilePath)
	if err != nil {
		log.Panic(err)
	}

	app := core.NewApplication(config, goMsServices, goMsMiddleWare)
	app.Start()
}
