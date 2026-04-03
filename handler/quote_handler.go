package handler

import (
	"brazilian-api-quotes/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type QuoteHandler struct {
	svc *service.QuoteService
}

func NewQuoteHandler(svc *service.QuoteService) *QuoteHandler {
	return &QuoteHandler{svc: svc}
}

func (h *QuoteHandler) GetQuotes(c *gin.Context) {
	pageStr := c.Query("page")
	pageSizeStr := c.Query("pageSize")

	if pageStr == "" && pageSizeStr == "" {
		quotes, err := h.svc.GetQuotes()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, quotes)
		return
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page"})
		return
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page size"})
		return
	}

	quotes, err := h.svc.GetPaginatedQuotes(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, quotes)
}
