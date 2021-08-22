//go:generate mockgen -destination=../mocks/mock_repo.go -package=repo github.com/ozonva/ova-book-api/internals/repo Repo

package repo

import (
	"github.com/ozonva/ova-book-api/internals/entities/book"
)

type Repo interface {
	AddEntities(entities []book.Book) error
	ListEntities(limit, offset uint64) ([]book.Book, error)
	DescribeEntity(entityId uint64) (*book.Book, error)
}
