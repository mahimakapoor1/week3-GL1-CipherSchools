package database

import (
	"gin/basic/models"

	"github.com/jinzhu/gorm"
)

func GetBooks(db *gorm.DB) ([]models.Book,error){
books:=[]models.Book{}
query:=db.Select("books.*")
err:=query.Find(&books).Error
if err!=nil{
	return nil,err
}

return books,nil
}
func GetBooksByID(db *gorm.DB,id string) (*models.Book,error){
	book:=models.Book{}
	err:= db.Select("books.*").Group("book.id").Where("book.id= ?",id).First(&book).Error
	if err!=nil{
		return nil,err
	}
	
	return &book,nil
}
func DeleteBooksByID(db *gorm.DB,id string) error{
	var book models.Book
	err:=db.Where("id=?",id).Delete(&book).Error
	if err!=nil{
		return err
	}
	return nil
}
func UpdateBooksByID(db *gorm.DB,book *models.Book) error{
	err:=db.Save(book).Error
		if err!=nil{
		return err
	}
	
	
	return nil
}
func SaveBook(db *gorm.DB,book *models.Book)error{
	err:=db.Save(book).Error
	if err!=nil{
		return err
	}
	
	return nil
}

