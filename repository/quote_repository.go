package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Quote struct {
	ID     int    `json:"id"`
	Quote  string `json:"quote"`
	Source string `json:"source"`
}

type QuoteRepository struct {
	db *pgxpool.Pool
}

func NewQuoteRepository(db *pgxpool.Pool) *QuoteRepository {
	return &QuoteRepository{db: db}
}

func (r *QuoteRepository) GetQuotes() ([]Quote, error) {
	quotes := []Quote{}

	rows, err := r.db.Query(context.Background(), "SELECT * FROM quotes;")
	if err != nil {
		return quotes, err
	}
	defer rows.Close()

	for rows.Next() {
		var quote Quote
		err := rows.Scan(&quote.ID, &quote.Quote, &quote.Source)
		if err != nil {
			return quotes, err
		}
		quotes = append(quotes, quote)
	}

	return quotes, nil
}
