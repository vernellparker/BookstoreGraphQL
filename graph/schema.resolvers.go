package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/vernellparker/BookstoreGraphQL/graph/generated"
	"github.com/vernellparker/BookstoreGraphQL/graph/model"
	inputHandler "github.com/vernellparker/BookstoreGraphQL/handlers/input"
)
// This resolver creates a new book
func (r *mutationResolver) CreateBook(ctx context.Context, input model.NewBook) (*model.Book, error) {
	ca := inputHandler.CheckForAuthorByNameInput{
		Name:     input.AuthorName,
		Authors:  r.AuthorsDB,
		BookISBN: input.Isbn,
	}

	authorID, err := inputHandler.CheckForAuthorByName(ca)
	if err != nil {
		return nil, err
	}
	book := &model.Book{
		Isbn:            input.Isbn,
		Title:           input.Title,
		PublicationDate: &input.PublishingDate,
		Stocked:         input.Stocked,
		Price:           input.Price,
		AuthorID:        authorID,
	}
	r.BooksDB = append(r.BooksDB, book)
	return book, nil
}
//Resolver to create an author
func (r *mutationResolver) CreateAuthor(ctx context.Context, input *model.NewAuthor) (*model.Author, error) {
	author := &model.Author{
		ID:      input.ID,
		Name:    input.Name,
		BookIds: nil,
	}
	r.AuthorsDB = append(r.AuthorsDB, author)
	return author, nil
}

//this resolver gets all books
func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	return r.BooksDB, nil
}

//this resolver gets all authors
func (r *queryResolver) Authors(ctx context.Context) ([]*model.Author, error) {
	return r.AuthorsDB, nil
}
// This resolver sorts the DB by price. This is just done to show how to take in input from a query and do something
// In practice you wouldn't have the logic sort the whole DB
func (r *queryResolver) SortBooksByPrice(ctx context.Context, ascending bool) ([]*model.Book, error) {
	books, err := inputHandler.SortBooks(r.BooksDB, ascending)
	if err != nil {
		return nil, err
	}
	return books, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

