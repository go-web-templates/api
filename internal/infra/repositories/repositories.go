package repositories

import (
	"github.com/go-web-templates/api/internal/application/interfaces/repositories"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	fx.Annotate(
		NewBooksRepository,
		fx.As(new(repositories.BooksRepository)),
	),
)
