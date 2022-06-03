package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Ayiruss/bookstore/graph/generated"
	"github.com/Ayiruss/bookstore/graph/internal/services"
	"github.com/Ayiruss/bookstore/graph/model"
	"github.com/Ayiruss/bookstore/graph/utils/service"
)

func (r *mutationResolver) CreateBook(ctx context.Context, input model.NewBook) (*model.Book, error) {
	return services.BookService.Create(ctx, &input)
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (interface{}, error) {
	return service.Login(ctx, input.UserName, input.Password)
}

func (r *queryResolver) ListBooks(ctx context.Context) ([]*model.Book, error) {
	return services.BookService.List()
}

func (r *mutationResolver) PurchaseBook(ctx context.Context, ID string) (*model.Book, error) {
	return services.BookService.Purchase(ctx, ID)
}

func (r *mutationResolver) ReSellBook(ctx context.Context, ID string) (*model.Book, error) {
	return services.BookService.Purchase(ctx, ID)
}

func (r *queryResolver) GetBook(ctx context.Context, ID string) (*model.Book, error) {
	return services.BookService.Get(ID)
}

func (r *queryResolver) Protected(ctx context.Context) (string, error) {
	return "Success", nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
