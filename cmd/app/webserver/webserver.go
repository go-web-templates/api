package webserver

import (
	"fmt"

	"github.com/go-web-templates/api/cmd/app/controllers"
	"github.com/go-web-templates/api/pkg/conf"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotate(
		NewFiberWebServer,
		fx.ParamTags(`group:"controllers"`),
		fx.As(new(WebServer)),
	),
)

type WebServer interface {
	StartServer() error
	ShutdownServer() error
}

type FiberWebServer struct {
	Server *fiber.App
	Config *conf.AppConf
}

func NewFiberWebServer(
	controllers []controllers.BaseController, // This must be the first parameter
	appConf *conf.AppConf,
) *FiberWebServer {
	server := fiber.New()

	server.Use(logger.New())

	for _, controller := range controllers {
		controller.RegisterController(server)
	}

	return &FiberWebServer{
		Server: server,
		Config: appConf,
	}
}

func (ws *FiberWebServer) StartServer() error {
	return ws.Server.Listen(fmt.Sprintf(":%s", ws.Config.Port))
}

func (ws *FiberWebServer) ShutdownServer() error {
	return ws.Server.Shutdown()
}
