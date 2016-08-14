package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Root(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to Elramir",
	})
}
