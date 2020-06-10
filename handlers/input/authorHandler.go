package input

import (
	"errors"
	"github.com/vernellparker/BookstoreGraphQL/graph/model"
)

type CheckForAuthorByNameInput struct {
	Name string
	Authors []*model.Author
	BookISBN string
}


//This function checks for the author by name and adds a the book to that author
func CheckForAuthorByName(input CheckForAuthorByNameInput) (string, error) {
	for _, author := range input.Authors {
		if input.Name == author.Name {
			author.BookIds = append(author.BookIds,input.BookISBN )
			return author.ID, nil
		}
	}
	return "", errors.New("author not in database")
}



