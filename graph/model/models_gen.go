// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Author struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	BookIds []string `json:"BookIds"`
}

type Book struct {
	Isbn            string  `json:"isbn"`
	Title           string  `json:"title"`
	PublicationDate *string `json:"publicationDate"`
	Stocked         bool    `json:"stocked"`
	Price           float64 `json:"price"`
	AuthorID        string  `json:"authorId"`
}

type NewAuthor struct {
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	BookIds []string `json:"BookIds"`
}

type NewBook struct {
	Isbn           string  `json:"isbn"`
	Title          string  `json:"title"`
	PublishingDate string  `json:"publishingDate"`
	Price          float64 `json:"price"`
	Stocked        bool    `json:"stocked"`
	AuthorID       string  `json:"authorId"`
}
