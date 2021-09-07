//go:generate mockgen -destination=./mock_repo.go -package=repo github.com/ozonva/ova-book-api/internals/repo Repo

package repo

import (
	"github.com/ozonva/ova-book-api/internals/entities/book"
)

type Repo interface {
	AddEntities(entities []book.Book) error
	ListEntities(limit, offset uint64) ([]book.Book, error)
	DescribeEntity(isbn10 string) (*book.Book, error)
	UpdateEntity(isbn10 string, newBook book.Book) (*book.Book, error)
	RemoveEntity(isbn10 string) error
}
