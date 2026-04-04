package main

import (
	"brazilian-api-quotes/config"
	"brazilian-api-quotes/di"
	"brazilian-api-quotes/handler"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbPool := config.LoadDatabase()
	defer dbPool.Close()

	svc := di.InitializeQuoteService(dbPool)
	h := handler.NewQuoteHandler(svc)
	engine := config.LoadEngine(h)

	engine.Run(":" + os.Getenv("PORT"))
}
