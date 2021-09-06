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

	booksEntities, err := a.Repo.ListEntities(2, 0)
	if err != nil {
		return nil, err
	}

	booksList := make([]*api.BookMessage, 0, 2)
	for _, book := range booksEntities {
		booksList = append(booksList, &api.BookMessage{
			UserId: book.UserId,
			Title:  book.Title,
			Author: book.Author,
			ISBN10: book.ISBN10,
			ISBN13: book.ISBN13,
		})
	}
	return &api.ListBooksMessage{
		Books: booksList,
	}, nil
}
