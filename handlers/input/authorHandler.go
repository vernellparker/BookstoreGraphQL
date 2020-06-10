package input

import (

	"github.com/vernellparker/BookstoreGraphQL/graph/model"
	"math/rand"
	"time"
)

func CheckForAuthorByName(name string, r []*model.Author) *model.Author {
	for _, author := range r {
		if name == author.Name {
			return author
		}
	}
	return CreateAnAuthor(name, nil)
}

func CreateAnAuthor(name string, books []model.Book) *model.Author{
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	author := &model.Author{
		ID:   string(r1.Intn(1000)),
		Name: name,
		Books: nil,
	}
	if books != nil {
		author.Books = books
	}
}