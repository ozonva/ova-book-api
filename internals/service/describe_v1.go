package service

import (
	"context"
	"fmt"

	"github.com/ozonva/ova-book-api/pkg/api"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func (a *BookApi) DescribeBook(ctx context.Context, book *api.CurrentBookMessage) (*api.BookMessage, error) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg(fmt.Sprintf("DescribeBook %v", book))

	currentBook, err := a.Repo.DescribeEntity(book.ISBN10)
	if err != nil {
		log.Info().Msg(fmt.Sprintf("При получении записи с ID 9 возикла ошибка: %v", err))
		return nil, err
	}

	return &api.BookMessage{
		UserId: currentBook.UserId,
		Title:  currentBook.Title,
		Author: currentBook.Author,
		ISBN10: currentBook.ISBN10,
		ISBN13: currentBook.ISBN13,
	}, nil
}
