package models

import (
	"github.com/jinzhu/gorm"
	"github.com/sheeeng/borneo-books/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	//Id          string `json:"id"`
	Name      string `gorm:""json:"name"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
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

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID = ?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID = ?", ID).Delete(book)
	return book
}
