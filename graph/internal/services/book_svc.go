package services

import (
	"context"
	"net/http"
	"time"

	"github.com/Ayiruss/bookstore/graph/internal/domain/books"
	"github.com/Ayiruss/bookstore/graph/internal/domain/users"
	"github.com/Ayiruss/bookstore/graph/internal/helper"
	"github.com/Ayiruss/bookstore/graph/model"
	"github.com/Ayiruss/bookstore/graph/utils/errors"
	"github.com/Ayiruss/bookstore/graph/utils/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	BookService bookServiceInterface = &bookService{}
)

type bookService struct{}

type bookServiceInterface interface {
	Create(context.Context, *model.NewBook) (*model.Book, error)
	List() ([]*model.Book, error)
	Get(string) (*model.Book, error)
	Purchase(context.Context, string) (*model.Book, error)
	ReSell(context.Context, string) (*model.Book, error)
}

func (s *bookService) Create(ctx context.Context, b *model.NewBook) (*model.Book, error) {
	claims := service.CtxValue(ctx)
	user := users.User{
		Username: claims.Username,
	} // Retrieving user from the claims that we use to fill up the additional details on BOOK DATA like SellerID
	err := user.GetByUserName()
	if err != nil {
		panic(err)
	}

	book := &books.Book{
		ID:          primitive.NewObjectID(),
		Title:       b.Title,
		Description: b.Description,
		Price:       helper.ToInt(b.Price),
		SellerID:    user.ID.Hex(),
		SellerName:  user.Name,
		Status:      helper.Published.EnumIndex(), // Published State
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		UpdatedBy:   user.ID.Hex(),
	}
	err = book.Create()
	return book.FromDB(), nil
}

func (s *bookService) List() ([]*model.Book, error) {
	books_do, err := books.List()
	if err != nil {
		return nil, &errors.MyError{
			Inner:      err,
			StatusCode: http.StatusInternalServerError,
			Message:    "Error reading data from the table",
		}
	}
	return books.FromDB(books_do), nil
}

func (s *bookService) Get(bookID string) (*model.Book, error) {
	objID, err := primitive.ObjectIDFromHex(bookID)
	if err != nil {
		return nil, &errors.MyError{
			Inner:      err,
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid ID",
		}
	}
	book := books.Book{ID: objID}
	if err := book.Get(); err != nil {
		return nil, &errors.MyError{
			Inner:      err,
			StatusCode: http.StatusInternalServerError,
			Message:    "Error reading data from the table",
		}
	}

	return book.FromDB(), nil
}

func (s *bookService) Purchase(ctx context.Context, bookID string) (*model.Book, error) {
	claims := service.CtxValue(ctx)
	user := users.User{
		Username: claims.Username,
	}
	err := user.GetByUserName()
	if err != nil {
		return nil, &errors.MyError{
			Inner:      err,
			StatusCode: http.StatusConflict,
			Message:    "Uhable to Authenticating the User",
		}
	}

	objID, err := primitive.ObjectIDFromHex(bookID)
	if err != nil {
		return nil, &errors.MyError{
			Inner:      err,
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid Book ID",
		}
	}
	book := books.Book{ID: objID}

	err = book.Get()

	if err != nil {
		return nil, &errors.MyError{
			Inner:      err,
			StatusCode: http.StatusBadRequest,
			Message:    "Book not found",
		}
	}

	if book.Price > user.Balance { // Comparing the user's balance with the price of the book
		return nil, &errors.MyError{
			Inner:      err,
			StatusCode: http.StatusConflict,
			Message:    "Not enough balance to make the purchase",
		}
	}
	if book.Status != helper.Unpublished.EnumIndex() {
		book.Status = helper.Unpublished.EnumIndex()
		book.UpdateStatus()
	}

	return book.FromDB(), nil
}

func (s *bookService) ReSell(ctx context.Context, bookID string) (*model.Book, error) {
	claims := service.CtxValue(ctx)
	user := users.User{
		Username: claims.Username,
	}
	err := user.GetByUserName()
	if err != nil {
		return nil, &errors.MyError{
			Inner:      err,
			StatusCode: http.StatusConflict,
			Message:    "Uhable to Authenticating the User",
		}
	}

	objID, err := primitive.ObjectIDFromHex(bookID)
	if err != nil {
		return nil, &errors.MyError{
			Inner:      err,
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid Book ID",
		}
	}
	book := books.Book{ID: objID}

	err = book.Get()
	if user.ID.Hex() != book.SellerID { // Verify the current user is the original seller of the book
		return nil, &errors.MyError{
			Inner:      err,
			StatusCode: http.StatusConflict,
			Message:    "Unable to Authenticate the user",
		}
	}

	if book.Status != helper.Published.EnumIndex() {
		book.Status = helper.Published.EnumIndex()
		book.UpdateStatus()
	}
	return book.FromDB(), nil
}
