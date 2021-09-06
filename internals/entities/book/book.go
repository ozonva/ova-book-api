package book

import (
	"fmt"
)

// Описывает сущность "Книга"
type Book struct {
	UserId uint64 `db:"user_id"`
	Title  string `db:"title"`
	Author string `db:"author"`
	ISBN10 string `db:"isbn10"`
	ISBN13 string `db:"isbn13"`
}

// Возврращает строковое представление для сущности "Книга"
func String(book *Book) string {
	return fmt.Sprintf("Book{UserID: %d, Title: %s}", book.UserId, book.Title)
}
