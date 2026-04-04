package main

import (
	"brazilian-api-quotes/config"
	"brazilian-api-quotes/di"
	"brazilian-api-quotes/handler"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	dbPool := config.LoadDatabase()
	defer dbPool.Close()

	svc := di.InitializeQuoteService(dbPool)
	h := handler.NewQuoteHandler(svc)
	engine := config.LoadEngine(h)

	engine.Run(":" + os.Getenv("PORT"))
}
