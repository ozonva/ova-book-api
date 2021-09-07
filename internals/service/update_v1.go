package service

import (
	"context"
	"fmt"

	"github.com/ozonva/ova-book-api/internals/entities/book"
	"github.com/ozonva/ova-book-api/pkg/api"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func (a *BookApi) UpdateBook(ctx context.Context, updateBook *api.BookMessage) (*api.BookMessage, error) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg(fmt.Sprintf("UpdateBook %v", updateBook))

	newBook := book.Book{
		UserId: updateBook.UserId,
		Title:  updateBook.Title,
		Author: updateBook.Author,
		ISBN10: updateBook.ISBN10,
		ISBN13: updateBook.ISBN13,
	}
	_, err := a.Repo.UpdateEntity(updateBook.ISBN10, newBook)
	if err != nil {
		log.Info().Msg(fmt.Sprintf("При добавлении записи возникла ошибка: %v", err))
		return nil, err
	}

	a.Producer.UpdateEvent(newBook)
	a.udpdatededCounter.Inc()
	return updateBook, nil
}
