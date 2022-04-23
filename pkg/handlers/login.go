package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *ChatAppController) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "ok")
}
