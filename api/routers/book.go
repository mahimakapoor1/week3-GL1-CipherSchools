package routers

import (
	"gin/basic/handler"

	"github.com/gin-gonic/gin"
)

func BookRouter(router *gin.Engine,api handler.Handler){
	// router := gin.Default()
	// api := handler.Handler{
	// 	DB:database.GetDB(),
	// }
	router.GET("/books", api.GetBooks)
	router.POST("/books",api.SaveBook)
	
}