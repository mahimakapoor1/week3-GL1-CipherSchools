package database

import (
	"errors"
	"gin/basic/models"

	"github.com/jinzhu/gorm"
)

func SignUp(db *gorm.DB,user *models.User)error{
	err:=db.Save(user).Error
	if(err!=nil){
		return err
	}
	return nil
}
func SignIn(db *gorm.DB,user *models.User) (error){
	getUser:=models.User{}
	err:= db.Select("users.*").Group("users.email_id").Where("users.email_id= ?",user.EmailId).First(&getUser).Error
	if err!=nil{
		return err
	}
	if user.Password!=getUser.Password{
		return errors.New("Password incorrect")
	}
	return nil
}