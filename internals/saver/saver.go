package saver

import (
	"errors"
	"fmt"
	"time"

	"github.com/ozonva/ova-book-api/internals/entities/book"
	"github.com/ozonva/ova-book-api/internals/flusher"
)

type Saver interface {
	Save(book book.Book) error
	Close() error
}

func timerFunc(ticker *time.Ticker, saver Saver) {
	for range ticker.C {
		err := saver.Close()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Данные были сохранены")
	}
}

type saver struct {
	books   []book.Book
	flusher flusher.Flusher
}

func (s *saver) Save(book book.Book) error {
	if len(s.books) == cap(s.books) {
		return errors.New("Достигнут лимит. Предварительно сохраните добавленные записи с помощью метода Close().")
	}

	s.books = append(s.books, book)
	return nil
}

func (s saver) Close() error {
	s.books = s.flusher.Flush(s.books)
	if len(s.books) != 0 {
		return errors.New("При сохранении возникла ошибка")
	}

	return nil
}

func New(capacity uint, flusher flusher.Flusher) Saver {
	saver_instannce := &saver{
		books:   make([]book.Book, 0, capacity),
		flusher: flusher,
	}
	ticker := time.NewTicker(time.Second * 5)
	go timerFunc(ticker, saver_instannce)

	return saver_instannce

}
