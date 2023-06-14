//go:generate mockgen -source=repository.go -destination=mocks/mock.go

package repository

import (
	"shortener/internal/repository/cache"
	"shortener/internal/repository/postgres"

	"github.com/jackc/pgx/v5"
)

type Repository interface {
	GetLongURL(short string) (string, error)
	AddLinksPair(short, long string) error
}

func NewRepositoryDB(conn *pgx.Conn) *postgres.Postgres {
	return postgres.New(conn)
}

func NewRepositoryIM() *cache.InMemory {
	return cache.New()
}
