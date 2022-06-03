package models

import (
	"github.com/jinzhu/gorm"
	"github.com/maitnngo2002/go_lang_projects/book-management-system/pkg/config"
)

var db *gorm.DB

type Book struct{
	gorm.Model
	Name string `gorm: "" json: "name"`
	Author string `json: "author"`
	Publication string `json: "publication"`

}

func init() {
	config.Connect() // connect to the database

	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
} 

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books 
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var returnedBook Book

	db := db.Where("ID=?", Id).Find(&returnedBook)

	return &returnedBook, db
}

func DeleteBook(Id int64) Book {
	var deletedBook Book

	db.Where("ID=?", Id).Delete(deletedBook)

	return deletedBook
}