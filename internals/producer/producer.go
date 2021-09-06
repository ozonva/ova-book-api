package producer

import "github.com/ozonva/ova-book-api/internals/entities/book"

type Producer interface {
	CreateEvent(createdBook book.Book) error
	UpdateEvent(updatedBook book.Book) error
	DeleteEvent(deletedBook book.Book) error
}
