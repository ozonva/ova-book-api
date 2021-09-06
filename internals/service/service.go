package service

import (
	"github.com/ozonva/ova-book-api/internals/counter"
	"github.com/ozonva/ova-book-api/internals/producer"
	"github.com/ozonva/ova-book-api/internals/repo"
	"github.com/ozonva/ova-book-api/pkg/api"
)

type BookApi struct {
	api.UnimplementedBookServiceServer

	Repo     repo.Repo
	Producer producer.Producer

	createdCounter    counter.Counter
	udpdatededCounter counter.Counter
	deletedCounter    counter.Counter
}

func NewBookApi(
	repo repo.Repo,
	producer producer.Producer,
	createdCounter, udpdatededCounter, deletedCounter counter.Counter,
) api.BookServiceServer {
	return &BookApi{
		Repo:              repo,
		Producer:          producer,
		createdCounter:    createdCounter,
		udpdatededCounter: udpdatededCounter,
		deletedCounter:    deletedCounter,
	}
}
