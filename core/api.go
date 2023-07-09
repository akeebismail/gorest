package core

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorest/controllers"
)

func Setup(ctx context.Context, router *gin.Engine) *gin.Engine {
	api := controllers.NewApiController(ctx)
	router.GET("/", api.Home)
	return router
}
