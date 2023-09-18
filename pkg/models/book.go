package models

import (
	"github.com/jinzhu/gorm"
	"github.com/nikhileshsirohi/BOOK-MANAGEMENT-SYSTEM/pkg/config"
)

var db *gorm.DB

type Book struct {
	//The error message you're seeing is because there is a syntax issue in your struct tags.
	//Struct tags in Go should not have spaces before or after the colons (':') or equal signs (=).
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBook() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookByID(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("id=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("id=?", Id).Delete(&book)
	return book
}
