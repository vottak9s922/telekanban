package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/telekanban/internal/application/usecases"
	"github.com/telekanban/internal/domain"
)

type BoardHandler struct {
	UC *usecases.BoardUsecase
}

func NewBoardHandler(uc *usecases.BoardUsecase) *BoardHandler {
	return &BoardHandler{UC: uc}
}

type CreateBoardInput struct {
	Name    string `json:"name" binding:"required"`
	OwnerID string `json:"owner_id" binding:"required"`
}

func (h *BoardHandler) Create(c *gin.Context) {
	var input CreateBoardInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	board, err := h.UC.CreateBoard(c, input.Name, input.OwnerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, board)
}

func (h *BoardHandler) List(c *gin.Context) {
	ownerID := c.Query("owner_id")
	if ownerID == "" {
		ownerID = "bob123"
	}
	boards, err := h.UC.BoardRepo.FindByOwner(c, ownerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if boards == nil {
		boards = make([]*domain.Board, 0)
	}
	c.JSON(http.StatusOK, boards)
}
