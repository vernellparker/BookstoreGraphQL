package graph

import "github.com/vernellparker/BookstoreGraphQL/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	BooksStorage []*model.Book
	Authors []*model.Author
}
