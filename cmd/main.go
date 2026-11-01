package main

import (
	"github.com/michelsazevedo/pagewise/internal/http"
	"github.com/michelsazevedo/pagewise/internal/http/logging"
	"github.com/michelsazevedo/pagewise/internal/infrastructure/database"
	"github.com/michelsazevedo/pagewise/internal/infrastructure/observability"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.WithLogger(logging.NewZerologfx),
		database.Db(),
		observability.Telemetry(),
		http.Http(),
	)

	app.Run()
}
