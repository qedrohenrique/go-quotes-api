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

func (s *QuoteService) GetPaginatedQuotes(page int, pageSize int) ([]repository.Quote, error) {

	if page < 1 {
		page = 1
	}

	if pageSize < 1 {
		pageSize = 10
	}

	quotes, err := s.repo.GetPaginatedQuotes(page, pageSize)
	if err != nil {
		return quotes, err
	}

	return quotes, nil
}