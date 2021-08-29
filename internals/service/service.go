package service

import "github.com/ozonva/ova-book-api/pkg/api"

type BookApi struct {
	api.UnimplementedBookServiceServer
}

func NewBookApi() api.BookServiceServer {
	return &BookApi{}
}
