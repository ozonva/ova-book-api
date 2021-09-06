package service

import (
	"context"
	"fmt"

	"github.com/ozonva/ova-book-api/internals/entities/book"
	"github.com/ozonva/ova-book-api/pkg/api"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func (a *BookApi) CreateBook(ctx context.Context, addBook *api.BookMessage) (*api.BookMessage, error) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg(fmt.Sprintf("CreatedBook %v", addBook))

	books := []book.Book{
		book.Book{
			UserId: addBook.UserId,
			Title:  addBook.Title,
			Author: addBook.Author,
			ISBN10: addBook.ISBN10,
			ISBN13: addBook.ISBN13,
		},
	}
	err := a.Repo.AddEntities(books)
	if err != nil {
		log.Info().Msg(fmt.Sprintf("При добавлении записи возникла ошибка: %v", err))
		return nil, err
	}

	return addBook, nil
}
