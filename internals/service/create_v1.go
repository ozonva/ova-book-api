package service

import (
	"context"
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/ozonva/ova-book-api/internals/entities/book"
	"github.com/ozonva/ova-book-api/internals/repo"
	"github.com/ozonva/ova-book-api/internals/utils"
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

	a.Producer.CreateEvent(books[0])
	a.createdCounter.Inc()
	return addBook, nil
}

func (a *BookApi) MultiCreateBook(ctx context.Context, booksList *api.ListBooksMessage) (*api.ListBooksMessage, error) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg(fmt.Sprintf("MultiCreatedBook %d", len(booksList.Books)))

	parentSpan, ctx := opentracing.StartSpanFromContext(ctx, "ParentMultiCreateBookSpan")
	defer parentSpan.Finish()

	books := make([]book.Book, 0, len(booksList.Books))

	for _, addBook := range booksList.Books {
		books = append(books,
			book.Book{
				UserId: addBook.UserId,
				Title:  addBook.Title,
				Author: addBook.Author,
				ISBN10: addBook.ISBN10,
				ISBN13: addBook.ISBN13,
			},
		)
	}

	for _, booksBatch := range utils.MakeBatchedSlice(books, 4) {
		err := createBatchOfBooks(ctx, a.Repo, booksBatch)
		if err != nil {
			log.Info().Msg(fmt.Sprintf("При добавлении записи возникла ошибка: %v", err))
			return nil, err
		}
		for _, createdBook := range books {
			a.Producer.CreateEvent(createdBook)
			a.createdCounter.Inc()
		}
	}

	return booksList, nil
}

func createBatchOfBooks(ctx context.Context, repo repo.Repo, booksBatch []book.Book) error {
	childSpan := opentracing.GlobalTracer().StartSpan(
		"ParentMultiCreateBookSpan",
		opentracing.ChildOf(opentracing.SpanFromContext(ctx).Context()))
	defer childSpan.Finish()

	err := repo.AddEntities(booksBatch)
	return err
}
