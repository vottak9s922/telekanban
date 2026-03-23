package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/telekanban/internal/application/usecases"
	"github.com/telekanban/internal/domain"
)

type ListHandler struct {
	UC *usecases.ListUsecase
}

func NewListHandler(uc *usecases.ListUsecase) *ListHandler {
	return &ListHandler{UC: uc}
}

type CreateListInput struct {
	Name     string `json:"name" binding:"required"`
	BoardID  string `json:"board_id" binding:"required"`
	Position int    `json:"position" binding:"required"`
}

func (h *ListHandler) Create(c *gin.Context) {
	var input CreateListInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	list, err := h.UC.CreateList(c, input.Name, input.BoardID, input.Position)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, list)
}

func (h *ListHandler) List(c *gin.Context) {
	boardID := c.Query("board_id")
	if boardID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "BoardID should be specified"})
	}
	lists, err := h.UC.ListRepo.FindByBoard(c, boardID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if lists == nil {
		lists = make([]*domain.List, 0)
	}
	c.JSON(http.StatusOK, lists)
}
