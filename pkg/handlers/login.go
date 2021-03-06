package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ShariUllas/chat-app/pkg/models"
	"github.com/gin-gonic/gin"
)

func (c *ChatAppController) Login(ctx *gin.Context) {
	req := &models.LoginRequest{}
	err := json.NewDecoder(ctx.Request.Body).Decode(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "invalid request body")
		return
	}
	users, err := c.service.Login(ctx, req)
	if err != nil {
		log.Println("Login failed:", err)
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(http.StatusOK, struct {
		UserID int `json:"user_id"`
	}{UserID: users.ID})
}
