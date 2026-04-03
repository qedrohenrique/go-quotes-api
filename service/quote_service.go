package service

import (
	"brazilian-api-quotes/repository"
)

type QuoteService struct {
	repo *repository.QuoteRepository
}

func NewQuoteService(repo *repository.QuoteRepository) *QuoteService {
	return &QuoteService{repo: repo}
}

func (s *QuoteService) GetQuotes() ([]repository.Quote, error) {
	quotes, err := s.repo.GetQuotes()
	
	if err != nil {
		return quotes, err
	}

	return quotes, nil
}
