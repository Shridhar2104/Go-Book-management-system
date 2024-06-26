package models

import (
    "github.com/Shridhar2104/Go-Book-management-system/pkg/config"
    "github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
    gorm.Model
    Name        string `json:"name"`
    Author      string `json:"author"`
    Publication string `json:"publication"`
}

func init() {
    config.Connect()
    db = config.GetDB()
    db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() error {
    result := db.Create(&b)
    return result.Error
}

func GetAllBooks() ([]Book, error) {
    var Books []Book
    result := db.Find(&Books)
    return Books, result.Error
}

func GetBookById(Id int64) (*Book, error) {
    var Book Book
    result := db.First(&Book, Id)
    return &Book, result.Error
}

func DeleteBook(Id int64) error {
    var Book Book
    result := db.Delete(&Book, Id)
    return result.Error
}

// func UpdateBook(Id int64, book *Book) error {
//     var Book Book
//     db.First(&Book, Id)
//     db.Model(&Book).Updates(book)
//     return nil
// }
