package service

import (
	"context"
	"fmt"

	"github.com/ozonva/ova-book-api/pkg/api"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func (a *BookApi) RemoveBook(ctx context.Context, book *api.CurrentBookMessage) (*api.BookMessage, error) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg(fmt.Sprintf("RemoveBook %v", book))

	return &api.BookMessage{
		UserId: 42,
		Title:  "Removed book's Title",
		Author: "Removed book's Author",
		ISBN10: book.ISBN10,
		ISBN13: "1234567890123",
	}, nil
}
