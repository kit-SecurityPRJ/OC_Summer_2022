package handler

import "github.com/gin-gonic/gin"

type Router struct {
	Engine gin.Engine
}

func NewRouter() *Router {
	router := &Router{
		Engine: *gin.Default(),
	}

	// portal page
	router.Engine.GET("/", func(ctx *gin.Context) {

	})
	// login test page
	router.Engine.GET("/login", func(ctx *gin.Context) {
		loginhandler(ctx)
	})
	// sql injection test page
	datapage := router.Engine.Group("/data")
	{
		datapage.GET("/search", func(ctx *gin.Context) {
			searchhandler(ctx)
		})
	}
	return router
}

func (router *Router) Run() {
	router.Engine.Run(":8080")
}
