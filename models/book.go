package models

import (
	"fmt"

	"github.com/nilerajput91/assig5bookapi/config"
)

// GetAllBook from db
func GetAllBook(b *[]Book) (err error) {
	if err = config.DB.Find(b).Error; err != nil {
		return err
	}

	return nil
}

// AddNewBook to db
func AddNewBook(b *Book) (err error) {
	if err = config.DB.Create(b).Error; err != nil {
		return err
	}
	return nil
}

// GetOneBook from db
func GetOneBook(b *Book, id string) (err error) {
	if err := config.DB.Where("id=?", id).First(b).Error; err != nil {
		return err
	}
	return nil
}

// PutOneBook to db
func PutOneBook(b *Book, id string) (err error) {
	fmt.Println(b)
	config.DB.Save(b)
	return nil
}

// DeleteBook from db
func DeleteBook(b *Book, id string) (err error) {
	config.DB.Where("id=?", id).Delete(b)
	return nil
}
