package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/telekanban/internal/db"
	"github.com/telekanban/internal/infrastruction/persistence/postgres"
	"go.uber.org/zap"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("no .env file found - using system vars")
	}

	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	sugar := logger.Sugar()

	if err := db.Connect(); err != nil {
		log.Fatalf("DB Connect failed: %v", err)
	}
	defer db.Pool.Close()

	boardRepo := postgres.NewBoardRepo(db.Pool)

}
