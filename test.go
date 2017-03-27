package main

import (
	_ "full_text_search/database"
	"full_text_search/controller"
)

func main() {
	controller.Test()

	//gin.SetMode(gin.DebugMode)
	//router := gin.Default()
	//router.GET("", controller.Test)
	//
	//router.Run(":9006")
}

