package http

import (
	"github.com/michelsazevedo/pagewise/internal/infrastructure/config"
	"go.uber.org/fx"
)

func Http() fx.Option {
	return fx.Module(
		"http",
		fx.Provide(
			config.NewConfig,
			NewBoot,
		),
		fx.Provide(
			fx.Annotate(NewHealthzHandler, fx.As(new(HealthzHandler))),
		),
		fx.Invoke(RegisterRoutes),
	)
}
