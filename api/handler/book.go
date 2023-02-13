package handler

import (
	"gin/basic/database"
	"gin/basic/models"
	"log"
	"net/http"
	// "github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"

)


func (h *Handler)GetBooks(in *gin.Context){
	books,err:=database.GetBooks(h.DB)
	if err!=nil{
		log.Println(err)
		in.JSON(500,err)
	}
	in.JSON(http.StatusOK,books)
}
func (h *Handler) SaveBook(in *gin.Context){
	book:=models.Book{}
	err:=in.BindJSON(&book)
	if err!=nil{
		log.Println(err)
		in.JSON(http.StatusInternalServerError,err)
	}
	err=database.SaveBook(h.DB,&book)
	if err!=nil{
		log.Println(err)
		in.JSON(http.StatusInternalServerError,err)
	}
	in.JSON(http.StatusOK,book)
}