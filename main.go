package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := NewServer()

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Working fine"})
	})

	router.Run(":8080")
}
