//go:build wireinject

package di

import (
	"brazilian-api-quotes/repository"
	"brazilian-api-quotes/service"

	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitializeQuoteService(db *pgxpool.Pool) *service.QuoteService {
	wire.Build(
		repository.NewQuoteRepository,
		service.NewQuoteService,
	)
	return nil
}
