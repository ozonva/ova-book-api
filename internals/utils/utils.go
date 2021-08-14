package utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ozonva/ova-book-api/internals/entities/book"
)

// Преобразует слайс в слайс слайсов указанной длины.
func MakeBatchedSlice(inSlice []book.Book, batchSize int) [][]book.Book {
	sliceSize := len(inSlice)
	batches := sliceSize / batchSize

	var resultSlice [][]book.Book
	startIdx := 0
	for idx := 0; idx < batches; idx++ {
		startIdx = idx * batchSize
		resultSlice = append(resultSlice, inSlice[startIdx:startIdx+batchSize])
	}
	nextIdx := startIdx + batchSize
	if nextIdx < sliceSize {
		resultSlice = append(resultSlice, inSlice[nextIdx:sliceSize])
	}
	return resultSlice
}

// Преобразует отображение ключ-значение в отображение значение-ключ.
func ReverseMap(inMap map[int]string) map[string]int {
	resultMap := map[string]int{}

	for key, value := range inMap {
		if _, ok := resultMap[value]; ok {
			panic("Ключ '" + value + "' уже существует")
		}
		resultMap[value] = key
	}
	return resultMap
}

// Фильтрует слайс по значениям, не фходящим в массив `numbers`.
func FilterIntSlice(inSlice []int) []int {
	var (
		numbers = [6]int{4, 8, 15, 16, 23, 42}
		result  []int
	)

	numbersMap := map[int]int{}
	for _, number := range numbers {
		numbersMap[number] = 1
	}

	for _, number := range inSlice {
		if _, ok := numbersMap[number]; !ok {
			result = append(result, number)
		}
	}
	return result
}

// Сравнивает на равенство две сущности "Книга"
func IsEqual(book, otherBook *book.Book) bool {
	return book.UserId == otherBook.UserId &&
		strings.EqualFold(book.ISBN10, otherBook.ISBN10) &&
		strings.EqualFold(book.ISBN13, otherBook.ISBN13)
}

// Преобразует слайс Книг в отображение.
func BookSliceToMap(books []book.Book) (map[uint64]book.Book, error) {
	result := map[uint64]book.Book{}

	for _, current_book := range books {
		if _, ok := result[current_book.UserId]; ok {
			return nil, errors.New(fmt.Sprintf("Для UserID: %d уже есть значение.", current_book.UserId))
		}
		result[current_book.UserId] = current_book
	}

	return result, nil
}
