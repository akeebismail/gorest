package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
)

type ApiController struct {
	ctx context.Context
}

func NewApiController(ctx context.Context) *ApiController {
	return &ApiController{
		ctx: ctx,
	}
}

func (api *ApiController) Home(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"data": "Hello world!",
	})
}
