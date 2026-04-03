package handler

import (
	"brazilian-api-quotes/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type QuoteHandler struct {
	svc *service.QuoteService
}

func NewQuoteHandler(svc *service.QuoteService) *QuoteHandler {
	return &QuoteHandler{svc: svc}
}

func (h *QuoteHandler) GetQuotes(c *gin.Context) {
	quotes, err := h.svc.GetQuotes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, quotes)
}