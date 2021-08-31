package saver_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ozonva/ova-book-api/internals/entities/book"
	"github.com/ozonva/ova-book-api/internals/flusher"
	"github.com/ozonva/ova-book-api/internals/repo"
	"github.com/ozonva/ova-book-api/internals/saver"
)

func getTestData() []book.Book {
	return []book.Book{
		{UserId: 1, Title: "First Title", Author: "First Author", ISBN10: "1234567891", ISBN13: "1234567890121"},
		{UserId: 2, Title: "Second Title", Author: "Second Author", ISBN10: "1234567892", ISBN13: "1234567890122"},
		{UserId: 3, Title: "Third Title", Author: "Third Author", ISBN10: "1234567893", ISBN13: "1234567890123"},
		{UserId: 4, Title: "Fourth Title", Author: "Fourth Author", ISBN10: "1234567894", ISBN13: "1234567890124"},
	}
}

func TestSaver(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	repoMock := repo.NewMockRepo(mockCtrl)
	books := getTestData()
	chunkSize := 4
	repoMock.EXPECT().AddEntities(books).Return(nil).AnyTimes()

	bookFlusher := flusher.New(uint(chunkSize), repoMock)
	bookSaver := saver.New(5, bookFlusher)
	for _, book := range books {
		bookSaver.Save(book)
	}

	err := bookSaver.Close()

	if err != nil {
		t.Error("Ожидается успешное сохранение записей.")
	}
}
