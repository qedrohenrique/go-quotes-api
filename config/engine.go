package config

import (
	"os"
	"strings"

	"brazilian-api-quotes/handler"

	"github.com/gin-gonic/gin"
)

func LoadEngine(h *handler.QuoteHandler) *gin.Engine {
	r := gin.Default()

	if proxies := os.Getenv("TRUSTED_PROXIES"); proxies != "" {
		r.SetTrustedProxies(strings.Split(proxies, ","))
	} else {
		r.SetTrustedProxies(nil)
	}

	r.GET("/quotes", h.GetQuotes)

	return r
}
