package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ShariUllas/chat-app/pkg/models"
	"github.com/gin-gonic/gin"
)

func (c *ChatAppController) SignUp(ctx *gin.Context) {
	req := &models.SignUpRequest{}
	err := json.NewDecoder(ctx.Request.Body).Decode(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "invalid request body")
		return
	}
	userID, err := c.service.SignUp(ctx, req)
	if err != nil {
		log.Println("Signup failed:", err)
		ctx.JSON(http.StatusInternalServerError, nil)
		return
	}
	ctx.JSON(http.StatusOK, struct {
		UserID int `json:"user_id"`
	}{UserID: userID})
}
