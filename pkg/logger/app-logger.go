package logger

import (
	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotate(
		NewTableLogger,
		fx.As(new(ApplicationLogger)),
	),
)

type ApplicationLogger interface {
	Warning(infos ...any)
	Info(infos ...any)
	Error(infos ...any)
	Fatal(infos ...any)
	Format(format string, data ...any) string
}
