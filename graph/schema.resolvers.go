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
//This resolver is for getting an author's books
func (r *authorResolver) Books(ctx context.Context, obj *model.Author) ([]*model.Book, error) {
	var books []*model.Book
	for _, b := range r.BooksDB{
		if b.Author.ID == obj.ID{
			books = append(books, b)
		}
	}
	return books, nil
}

//This resolver is for get a book's author
func (r *bookResolver) Author(ctx context.Context, obj *model.Book) (*model.Author, error) {
	for _, a := range r.AuthorsDB{
		if a.ID == obj.Author.ID{
			return a, nil
		}
	}
	return nil, errors.New("error finding author")
}

func (r *mutationResolver) CreateBook(ctx context.Context, input model.NewBook) (*model.Book, error) {
	ca := inputHandler.CheckForAuthorByNameInput{
		Name:     input.AuthorName,
		Authors:  r.AuthorsDB,
	}

	author, err := inputHandler.CheckForAuthorByName(ca)
	if err != nil {
		return nil, err
	}
	book := &model.Book{
		Isbn:            input.Isbn,
		Title:           input.Title,
		PublicationDate: &input.PublishingDate,
		Stocked:         input.Stocked,
		Price:           input.Price,
		Author: 		author,
	}
	r.BooksDB = append(r.BooksDB, book)
	return book, nil
}

func (r *mutationResolver) CreateAuthor(ctx context.Context, input *model.NewAuthor) (*model.Author, error) {
	author := &model.Author{
		ID:    input.ID,
		Name:  input.Name,
		Books: nil,
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

func (r *queryResolver) Author(ctx context.Context, id string) (*model.Author, error) {
	for _, a := range r.AuthorsDB {
		if a.ID == id {
			return a, nil
		}
	}
	return nil, errors.New("no author with that id was found")
}

func (r *queryResolver) Book(ctx context.Context, isbn string) (*model.Book, error) {
	for _, b := range r.BooksDB {
		if b.Isbn == isbn {
			return b, nil
		}
	}
	return nil, errors.New("no book with that isbn was found")
}

// Author returns generated.AuthorResolver implementation.
func (r *Resolver) Author() generated.AuthorResolver { return &authorResolver{r} }

// Book returns generated.BookResolver implementation.
func (r *Resolver) Book() generated.BookResolver { return &bookResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type authorResolver struct{ *Resolver }
type bookResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
