package flusher_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ozonva/ova-book-api/internals/entities/book"
	"github.com/ozonva/ova-book-api/internals/flusher"
	"github.com/ozonva/ova-book-api/internals/repo"
)

var _ = Describe("FlusherGinkgo", func() {
	var books []book.Book

	BeforeEach(func() {
		books = []book.Book{
			{UserId: 1, Title: "First Title", Author: "First Author", ISBN10: "1234567891", ISBN13: "1234567890121"},
			{UserId: 2, Title: "Second Title", Author: "Second Author", ISBN10: "1234567892", ISBN13: "1234567890122"},
			{UserId: 3, Title: "Third Title", Author: "Third Author", ISBN10: "1234567893", ISBN13: "1234567890123"},
			{UserId: 4, Title: "Fourth Title", Author: "Fourth Author", ISBN10: "1234567894", ISBN13: "1234567890124"},
		}

	})

	Describe("Сохранение записей", func() {
		Context("Без ошибок", func() {
			It("Должно быть равно nil", func() {
				mockCtrl := gomock.NewController(GinkgoT())
				defer mockCtrl.Finish()

				repoMock := repo.NewMockRepo(mockCtrl)
				chunkSize := 4
				repoMock.EXPECT().AddEntities(books).Return(nil)

				bookFlusher := flusher.New(uint(chunkSize), repoMock)
				expectedResult := bookFlusher.Flush(books) == nil
				Expect(expectedResult).To(Equal(true))
			})
		})
	})
})
