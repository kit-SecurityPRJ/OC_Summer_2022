package handler

import (
	"OCsemmerApp/pkg/model"
	"log"

	"github.com/gin-gonic/gin"
)

func searchhandler(ctx *gin.Context) {
	prodname := ctx.Query("prodname")
	if prodname == "" {
		log.Println("not prodname request")
		ctx.JSON(400, "Bad Request")
		return
	}
	data, err := model.SearchData(prodname)
	if err != nil {
		log.Println(err)
		ctx.JSON(500, "Internal Server Error")
		return
	}
	ctx.JSON(200, data)
}
