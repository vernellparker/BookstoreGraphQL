package input

import (
	"github.com/pkg/errors"
	"github.com/vernellparker/BookstoreGraphQL/graph/model"
)

type CheckForAuthorByNameInput struct {
	Name string
	Authors []*model.Author
}


//This function checks for the author by name and adds a the book to that author
func CheckForAuthorByName(input CheckForAuthorByNameInput) (*model.Author, error) {
	for _, author := range input.Authors {
		if input.Name == author.Name {
			//author.Books = append(author.Books,input.BookISBN )
			return author, nil
		}
	}
	return nil, errors.New("author not in database")
}



