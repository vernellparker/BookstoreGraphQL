package input

import (
	"errors"
	"github.com/vernellparker/BookstoreGraphQL/graph/model"
	)


//Sorts the books in the DB
func SortBooks(books []*model.Book, ascending bool) ([]*model.Book, error) {
	b:=len(books)
	if b == 0 {
		return nil, errors.New("no books in database")
	}
	//good old slow bubble sort
	//never use this algorithm in production
	if ascending == true {
		for i:=0;i<b;i++{
			for j:=0;j< (b-1-i);j++{
				if books[j].Price > books[j+1].Price{
					books[j], books[j+1] = books[j+1], books[j]
				}
			}
		}
	}else {
		for i:=0;i<b;i++{
			for j:=0;j< (b-1-i);j++{
				if books[j].Price < books[j+1].Price{
					books[j], books[j+1] = books[j+1], books[j]
				}
			}
		}
	}

	return books, nil
}