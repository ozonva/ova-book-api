package service

import (
	"context"
	"fmt"

	"github.com/ozonva/ova-book-api/pkg/api"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func (a *BookApi) RemoveBook(ctx context.Context, removeBook *api.CurrentBookMessage) (*api.BookMessage, error) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg(fmt.Sprintf("RemoveBook %v", removeBook))
	descrBook, err := a.Repo.DescribeEntity(removeBook.ISBN10)
	if err != nil {
		log.Info().Msg(fmt.Sprintf("Ошибка при получении записи: %v", err))
		return nil, err
	}
	err = a.Repo.RemoveEntity(removeBook.ISBN10)
	if err != nil {
		log.Info().Msg(fmt.Sprintf("Ошибка при удалении записи: %v", err))
		return nil, err
	}

	a.Producer.DeleteEvent(*descrBook)
	a.deletedCounter.Inc()
	return &api.BookMessage{
		UserId: descrBook.UserId,
		Title:  descrBook.Title,
		Author: descrBook.Author,
		ISBN10: descrBook.ISBN10,
		ISBN13: descrBook.ISBN13,
	}, nil
}
