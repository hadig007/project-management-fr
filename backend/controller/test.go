package controller

import "github.com/gin-gonic/gin"

func Test(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "server running",
	})
}
