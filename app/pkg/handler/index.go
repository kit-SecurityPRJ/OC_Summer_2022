package handler

import (
	"github.com/gin-gonic/gin"
)

func indexportalHandler(ctx *gin.Context) {
	ctx.HTML(200, "index.html", nil)
}
