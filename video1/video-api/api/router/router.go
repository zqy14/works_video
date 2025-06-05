package router

import (
	"github.com/gin-gonic/gin"
	"video1/video-api/api/handler"
)

func LandRouter(r *gin.Engine) {
	api := r.Group("/api")
	{
		user := api.Group("/user")
		{
			user.POST("/login", handler.Login)
		}

	}
}
