package flusher

import (
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/ozonva/ova-book-api/internals/entities/book"
	"github.com/ozonva/ova-book-api/internals/flusher"
	"github.com/ozonva/ova-book-api/internals/repo"
	"github.com/ozonva/ova-book-api/internals/utils"
)

func getTestData() ([]book.Book, []book.Book, []book.Book) {
	firstPart := []book.Book{
		{UserId: 1, Title: "First Title", Author: "First Author", ISBN10: "1234567891", ISBN13: "1234567890121"},
		{UserId: 2, Title: "Second Title", Author: "Second Author", ISBN10: "1234567892", ISBN13: "1234567890122"},
	}
	secondPart := []book.Book{
		{UserId: 3, Title: "Third Title", Author: "Third Author", ISBN10: "1234567893", ISBN13: "1234567890123"},
		{UserId: 4, Title: "Fourth Title", Author: "Fourth Author", ISBN10: "1234567894", ISBN13: "1234567890124"},
	}
	var books []book.Book
	books = append(books, firstPart...)
	books = append(books, secondPart...)
	return firstPart, secondPart, books
}

func TestPositiveFlusher(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	repoMock := repo.NewMockRepo(mockCtrl)
	_, _, books := getTestData()
	chunkSize := 4
	repoMock.EXPECT().AddEntities(books).Return(nil)

	bookFlusher := flusher.New(uint(chunkSize), repoMock)
	notFlushed := bookFlusher.Flush(books)
	if notFlushed != nil {
		t.Error("Ошибка при сбросе записей.")
	}
}

func TestNegativeFlusher(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	repoMock := repo.NewMockRepo(mockCtrl)
	_, _, books := getTestData()
	chunkSize := 4
	repoMock.EXPECT().AddEntities(books).Return(errors.New("Test"))

	bookFlusher := flusher.New(uint(chunkSize), repoMock)
	notFlushed := bookFlusher.Flush(books)
	for idx, book := range notFlushed {
		expectedBook := books[idx]
		if !utils.IsEqual(&book, &expectedBook) {
			t.Error(fmt.Sprintf("Не совпадает запись %d с %d", book.UserId, expectedBook.UserId))
		}
	}
}
