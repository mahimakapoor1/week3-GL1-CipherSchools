package main

import (
	"fmt"
	"gin/basic/database"
	"gin/basic/routers"

	"gin/basic/handler"
	"log"

	"github.com/gin-gonic/gin"
)
func init(){
	database.Setup()
}
func AuthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
	token:=c.GetHeader("token")
	fmt.Println("got token:"+token)
	isValid,err:=handler.ValidateToken(token)
	if err!=nil && !isValid{
		return 
	}
	c.Next()
}
}
func main(){
	// database.Setup()
	// engine:=routers.Router()
	engine := gin.Default()
	api := handler.Handler{
		DB:database.GetDB(),
	}
	routers.BookRouter(engine,api)
	routers.AuthRouter(engine,api)
	engine.Use(AuthMiddleware())
	err:=engine.Run("127.0.0.1:8080")
	if err!=nil{
		log.Fatal(err)
	}

}