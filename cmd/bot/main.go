package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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

	r := gin.Default()

	r.GET("/boards", func(c *gin.Context) {
		ownerID := c.Query("owner_id")
		if ownerID == "" {
			ownerID = "test_owner"
		}
		boards, err := boardRepo.FindByOwner(context.Background(), ownerID)
		if err != nil {
			sugar.Errorw("Failed to fetch boards", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		sugar.Infow("Boards fetched", "count", len(boards), "owner_id", ownerID)
		c.JSON(http.StatusOK, boards)
	})

}
