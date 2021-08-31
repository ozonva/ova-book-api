package service

import (
	"context"
	"fmt"

	"github.com/ozonva/ova-book-api/pkg/api"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func (a *BookApi) CreateBook(ctx context.Context, book *api.BookMessage) (*api.BookMessage, error) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg(fmt.Sprintf("CreatedBook %v", book))

	return &api.BookMessage{
		UserId: 123,
		Title:  "Title",
		Author: "Author",
		ISBN10: "1234567890",
		ISBN13: "1234567890123",
	}, nil
}
