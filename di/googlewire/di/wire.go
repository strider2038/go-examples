//go:build wireinject

package di

import (
	"net/http"

	"googlewire/app/db"
	"googlewire/app/domain"

	"github.com/google/wire"
)

func InitializeServer() (*http.Server, error) {
	wire.Build(
		NewServer,
		NewFindEntityHandler,
		NewRequestTimeout,
		NewFindEntityUseCase,
		NewEntityRepository,
		NewDB,
		NewConfig,
		wire.Bind(new(domain.EntityRepository), new(*db.EntityRepository)),
	)

	return &http.Server{}, nil
}
