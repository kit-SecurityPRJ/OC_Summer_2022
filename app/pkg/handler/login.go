package handler

import (
	"OCsemmerApp/pkg/domain"
	"OCsemmerApp/pkg/model"
	"log"

	"github.com/gin-gonic/gin"
)

func loginhandler(ctx *gin.Context) {
	login := domain.LoginRequest{}
	if err := ctx.Bind(&login); err != nil {
		log.Println(err)
		ctx.JSON(400, "Bad Request")
		return
	}
	if err := model.SearchUser(&login); err != nil {
		log.Println(err)
		// エラー文をバックエンド側の文を送ってしまうやつ
		ctx.JSON(500, err.Error())
		return
	}
	ctx.JSON(200, "OK")
}
