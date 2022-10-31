package api

import (
	"context"
	"fx-test/api/config"
	"fx-test/api/handler/echo"
	"fx-test/api/handler/hello"
	"fx-test/api/server"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func API() []fx.Option {
	return []fx.Option{
		fx.WithLogger(func(lgr *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: lgr}
		}),
		fx.Provide(
			zap.NewExample,
			config.InitializeAPIConfig,
			asRoute(echo.NewEchoHandler),
			asRoute(hello.NewHelloHandler),
			server.NewHTTPServer,
			fx.Annotate(
				server.NewServeMux,
				fx.ParamTags(`group:"routes"`),
			),
		),
		fx.Invoke(
			startServer,
		),
	}
}

var startServer = func(lc fx.Lifecycle, srv *server.Server) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return srv.StartServer()
		},
		OnStop: func(ctx context.Context) error {
			return srv.StopServer(ctx)
		},
	})
}

func asRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(server.Route)),
		fx.ResultTags(`group:"routes"`),
	)
}
