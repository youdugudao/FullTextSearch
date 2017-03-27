package main

import (
	_ "full_text_search/config"
	"github.com/gin-gonic/gin"
	"full_text_search/controller"
	"full_text_search/search"
)

func main() {
	search.InitSeg()
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.GET("", controller.Get)
	router.POST("", controller.Post)
	router.PUT("", controller.Put)
	router.DELETE("", controller.Delete)

	router.Run(":9005")
}
