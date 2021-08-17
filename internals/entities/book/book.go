package book

import (
	"fmt"
)

// Описывает сущность "Книга"
type Book struct {
	UserId uint64
	Title  string
	Author string
	ISBN10 string
	ISBN13 string
}

// Возврращает строковое представление для сущности "Книга"
func String(book *Book) string {
	return fmt.Sprintf("Book{UserID: %d, Title: %s}", book.UserId, book.Title)
}
