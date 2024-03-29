package main

import (
	"flag"
	"log"

	"github.com/reversTeam/go-ms-tools/middlewares"
	"github.com/reversTeam/go-ms-tools/services/account"
	"github.com/reversTeam/go-ms-tools/services/email"
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

var goMsServices = map[string]core.GoMsServiceFunc{
	"people": core.RegisterServiceMap(func(ctx *core.Context, name string, config core.ServiceConfig) core.GoMsServiceInterface {
		return people.NewService(ctx, name, config)
	}),
	"signin": core.RegisterServiceMap(func(ctx *core.Context, name string, config core.ServiceConfig) core.GoMsServiceInterface {
		return signin.NewService(ctx, name, config)
	}),
	"account": core.RegisterServiceMap(func(ctx *core.Context, name string, config core.ServiceConfig) core.GoMsServiceInterface {
		return account.NewService(ctx, name, config)
	}),
	"email": core.RegisterServiceMap(func(ctx *core.Context, name string, config core.ServiceConfig) core.GoMsServiceInterface {
		return email.NewService(ctx, name, config)
	}),
}
var goMsMiddleWare = map[string]core.Middleware{
	"CheckParametersMiddleware": middlewares.NewCheckParametersMiddleware(),
	"AuthMiddleware":            &middlewares.AuthMiddleware{},
	"UnAuthMiddleware":          &middlewares.UnauthenticatedMiddleware{},
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
