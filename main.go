package main

import (
	"chat-app/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// http.Handle("/", http.FileServer(http.Dir("./static")))
	// Creates a router without any middleware by default
	router := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	router.StaticFS("/index", http.Dir("./static"))

	v1 := router.Group("/")
	{
		v1.GET("/chat", server.WSEndPoint)
	}

	router.Run(":8080")
}
