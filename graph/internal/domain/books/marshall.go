package books

import (
	"github.com/Ayiruss/bookstore/graph/internal/helper"
	"github.com/Ayiruss/bookstore/graph/model"
)

func (book *Book) FromDB() *model.Book {

	result := &model.Book{
		ID:          book.ID.Hex(),
		Title:       book.Title,
		Description: &book.Description,
		Price:       helper.ToFloat(book.Price),
		SellerName:  &book.SellerName,
		Status:      helper.Status(book.Status).String(),
	}

	return result
}

func FromDB(books_do []*Book) []*model.Book {
	var books []*model.Book
	for _, book_do := range books_do {
		book := book_do.FromDB()
		books = append(books, book)
	}
	return books
}
