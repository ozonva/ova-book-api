package service

import (
	"github.com/ozonva/ova-book-api/internals/repo"
	"github.com/ozonva/ova-book-api/pkg/api"
)

type BookApi struct {
	api.UnimplementedBookServiceServer

	Repo repo.Repo
}

func NewBookApi(repo repo.Repo) api.BookServiceServer {
	return &BookApi{Repo: repo}
}
