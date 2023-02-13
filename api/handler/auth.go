package handler

import (
	"gin/basic/database"
	"gin/basic/models"
	"log"
	"net/http"
	
	"github.com/gin-gonic/gin"
	jwt "get github.com/dgrijalva/jwt-go"
)

func (h *Handler) SignUp(in *gin.Context) {
	user := models.User{}
	err := in.BindJSON(&user)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	err = database.SignUp(h.DB, &user)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	in.JSON(http.StatusOK, user)
}
func (h *Handler) SignIn(in *gin.Context) {
	user := models.User{}
	err := in.BindJSON(&user)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	err = database.SignIn(h.DB, &user)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	token, err:=generateToken(user)
	if err != nil {
		log.Println(err)
		in.JSON(http.StatusInternalServerError, err)
	}
	in.JSON(http.StatusOK, token)
}
func generateToken(user models.User)(string,error){
	token:=	jwt.New(jwt.SigningMethodH5256)
	claims:=token.Claims.(jwt.MapClaims)
	claims["authorized"]=true
	claims["emailed"]=user.EmailId
	claims["exp"]=time.Now().Add(time.Minute*30).Unix()
	tokenString:=token.SignedString("rahul")
	if err!=nil{
		log.Println(err)
		return "",err
	}
	return tokenString,nil
}
func ValidateToken(token string)(bool,error){
	jwt.Parsel(Intoken,func(token *jwt.Token)(interface{},error){
		_,ok:=token.Method.(*jwt.SigningMethodHMAC)
		if !ok{
			return nil,errors.New("some unexpected error")
		}
		return []byte("rahul"),nil
	})
	if err != nil {
		log.Println(err)
		return false,err
	}
	if !token.valid {
		return false,errors.New("token is not valid")
	}
	return true,nil
}