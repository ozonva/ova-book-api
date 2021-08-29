package service

import (
	"context"

	"github.com/ozonva/ova-book-api/pkg/api"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *BookApi) ListBooks(ctx context.Context, _ *emptypb.Empty) (*api.ListBooksMessage, error) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("ListBooks")

	books := []*api.BookMessage{
		&api.BookMessage{
			UserId: 321,
			Title:  "Describe",
			Author: "Describe Author",
			ISBN10: "1234567890",
			ISBN13: "1234567890123",
		},
		&api.BookMessage{
			UserId: 123,
			Title:  "Title",
			Author: "Author",
			ISBN10: "1234567890",
			ISBN13: "1234567890123",
		},
	}
	return &api.ListBooksMessage{
		Books: books,
	}, nil
}
