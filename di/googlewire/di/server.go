package di

import (
	"database/sql"
	"net/http"
	"time"

	"googlewire/app/db"
	"googlewire/app/domain"
	"googlewire/app/httphandler"
	"googlewire/app/usecase"
)

func NewServer(
	findEntityHandler http.Handler,
	requestTimeout time.Duration,
) *http.Server {
	return &http.Server{
		Handler:     findEntityHandler,
		IdleTimeout: requestTimeout,
	}
}

func NewFindEntityHandler(u *usecase.FindEntity) http.Handler {
	return httphandler.NewFindEntity(u)
}

func NewFindEntityUseCase(entities domain.EntityRepository) *usecase.FindEntity {
	return usecase.NewFindEntity(entities)
}

func NewEntityRepository(conn *sql.DB) *db.EntityRepository {
	return db.NewEntityRepository(conn)
}

func NewDB(config *Config) (*sql.DB, error) {
	return sql.Open("postgres", config.DatabaseURL)
}

func NewRequestTimeout() time.Duration {
	return time.Second
}
