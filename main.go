package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorest/core"
)

func main() {
	ctx := context.Background()
	router := gin.Default()
	core.Setup(ctx, router)
	router.Run(":3000")
}
