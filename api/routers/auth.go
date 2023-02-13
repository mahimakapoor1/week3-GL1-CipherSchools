package routers

import (
	
	"gin/basic/handler"
	"github.com/gin-gonic/gin"
)

func AuthRouter(router *gin.Engine,api handler.Handler) {
	router.POST("/signup",api.SignUp)
	router.POST("/signin",api.SignIn)

}