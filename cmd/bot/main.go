package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/telekanban/internal/application/usecases"
	"github.com/telekanban/internal/db"
	"github.com/telekanban/internal/handlers"
	"github.com/telekanban/internal/infrastructure/persistence/postgres"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("no .env file found - using system vars")
	}

	if err := db.Connect(); err != nil {
		log.Fatalf("DB Connect failed: %v", err)
	}
	defer db.Pool.Close()

	boardRepo := postgres.NewBoardRepo(db.Pool)
	boardUC := usecases.NewBoardUsecase(boardRepo)
	boardHandler := handlers.NewBoardHandler(boardUC)

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/boards", boardHandler.List)
	r.POST("/boards", boardHandler.Create)

	r.Run(":8080")
}
