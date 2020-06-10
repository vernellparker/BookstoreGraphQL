package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"github.com/vernellparker/BookstoreGraphQL/graph/generated"
	"github.com/vernellparker/BookstoreGraphQL/graph/model"
	inputHandler "github.com/vernellparker/BookstoreGraphQL/handlers/input"
)

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

func (r *mutationResolver) CreateAuthor(ctx context.Context, input *model.NewAuthor) (*model.Author, error) {
	author := &model.Author{
		ID:      input.ID,
		Name:    input.Name,
		BookIds: nil,
	}
	r.AuthorsDB = append(r.AuthorsDB, author)
	return author, nil
}

func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	return r.BooksDB, nil
}

func (r *queryResolver) Authors(ctx context.Context) ([]*model.Author, error) {
	return r.AuthorsDB, nil
}

func (r *queryResolver) SortBooksByPrice(ctx context.Context, ascending bool) ([]*model.Book, error) {
	books, err := inputHandler.SortBooks(r.BooksDB, ascending)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (r *queryResolver) GetAuthor(ctx context.Context, id string) (*model.Author, error) {
	for _, a := range r.AuthorsDB {
		if a.ID == id {
			return a, nil
		}
	}
	return nil, errors.New("no author with that id was found")
}

func (r *queryResolver) GetBook(ctx context.Context, isbn string) (*model.Book, error) {
	for _, b := range r.BooksDB {
		if b.Isbn == isbn {
			return b, nil
		}
	}
	return nil, errors.New("no book with that isbn was found")
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
