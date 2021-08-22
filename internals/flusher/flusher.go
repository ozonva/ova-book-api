package flusher

import (
	"github.com/ozonva/ova-book-api/internals/entities/book"
	"github.com/ozonva/ova-book-api/internals/repo"
	"github.com/ozonva/ova-book-api/internals/utils"
)

type Flusher interface {
	Flush(entities []book.Book) []book.Book
}

type flusher struct {
	chunkSize  uint
	entityRepo repo.Repo
}

func New(chunkSize uint, entitiyRepo repo.Repo) Flusher {
	return &flusher{
		chunkSize:  chunkSize,
		entityRepo: entitiyRepo,
	}
}

func (f flusher) Flush(entities []book.Book) []book.Book {
	var notSavedEntites []book.Book
	for _, batch := range utils.MakeBatchedSlice(entities, int(f.chunkSize)) {
		err := f.entityRepo.AddEntities(batch)
		if err != nil {
			notSavedEntites = append(notSavedEntites, batch...)
		}
	}
	return notSavedEntites
}
