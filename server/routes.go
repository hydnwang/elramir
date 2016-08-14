package server

import (
	"github.com/gin-gonic/gin"
	"github.com/hydnwang/elramir/config"
	"github.com/hydnwang/elramir/handlers"
)

func RoutersEngine() *gin.Engine {
	// set server mode
	gin.SetMode(config.Mode)

	r := gin.Default()

	r.GET("/", handlers.Root)

	v1 := r.Group("api/v1")
	{
		v1.GET("/photos/:cid", handlers.GetPhotos)
		v1.POST("/photos", handlers.PostPhoto)
		v1.GET("/photos/:cid/:id", handlers.GetPhoto)
		v1.GET("/thumb/:cid/:id", handlers.GetThumbnail)
	}

	return r
}
