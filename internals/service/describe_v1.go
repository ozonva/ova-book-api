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

	return &api.BookMessage{
		UserId: 321,
		Title:  "Describe",
		Author: "Describe Author",
		ISBN10: "1234567890",
		ISBN13: "1234567890123",
	}, nil
}
