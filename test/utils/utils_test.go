package utils

import (
	"fmt"
	"testing"

	"github.com/ozonva/ova-book-api/internals/entities/book"
	"github.com/ozonva/ova-book-api/internals/utils"
)

type testBatchedSliceCases struct {
	inSlice        []book.Book
	batchSize      int
	expectedSlices [][]book.Book
}

func generateBookSlice() []book.Book {
	var books []book.Book
	for idx := 0; idx < 6; idx++ {
		books = append(books, book.Book{
			UserId: uint64(idx),
			Title:  fmt.Sprintf("Test %d", idx),
			Author: fmt.Sprintf("Author %d", idx),
			ISBN10: "1234567890",
			ISBN13: "1234567890123",
		})
	}
	return books
}

func prepareMakeBatchedSliceData() []testBatchedSliceCases {
	books := generateBookSlice()
	return []testBatchedSliceCases{
		{
			books[:6], 2, [][]book.Book{{books[0], books[1]}, {books[2], books[3]}, {books[4]}},
		},
		{
			books[:5], 3, [][]book.Book{{books[0], books[1], books[2]}, {books[3], books[4]}},
		},
		{
			books, 3, [][]book.Book{{books[0], books[1], books[2]}, {books[3], books[4], books[5]}},
		},
	}
}

func TestMakeBatchedSlice(t *testing.T) {
	var testCases []testBatchedSliceCases = prepareMakeBatchedSliceData()

	for _, testCase := range testCases {
		expected := testCase.expectedSlices
		inSlice := testCase.inSlice
		batchSize := testCase.batchSize

		outSlice := utils.MakeBatchedSlice(inSlice, batchSize)

		if len(expected) != len(outSlice) {
			t.Error("Не совпадает длина слайса, ожидается: ", len(expected), "Получено: ", len(outSlice))
		} else {
			for idx, value := range expected {
				for innerIdx, innerValue := range value {
					if !utils.IsEqual(&innerValue, &outSlice[idx][innerIdx]) {
						t.Error(value, "!=", outSlice[idx])
					}
				}
			}
		}
	}
}

type reverseMapTestCase struct {
	inMap       map[int]string
	expectedMap map[string]int
}

func preparePositiveReverseMapTestCases() []reverseMapTestCase {
	return []reverseMapTestCase{
		{map[int]string{1: "one", 2: "two"}, map[string]int{"one": 1, "two": 2}},
		{map[int]string{3: "three", 4: "four"}, map[string]int{"three": 3, "four": 4}},
	}
}

func TestReverseMap(t *testing.T) {
	testCases := preparePositiveReverseMapTestCases()

	for _, testCase := range testCases {
		caseMap := testCase.inMap
		expectedMap := testCase.expectedMap
		resultMap := utils.ReverseMap(caseMap)
		for key := range expectedMap {
			if _, ok := resultMap[key]; !ok {
				t.Error("Не найден ключ:", key, "в полученном отображении:", resultMap)
			}
		}
	}
}

func TestPanicReverseMap(t *testing.T) {

	defer func() {
		if r := recover(); r == nil {
			t.Error("Не была вызвана panic()")
		}
	}()

	inMap := map[int]string{1: "one", 2: "two", 3: "two"}
	utils.ReverseMap(inMap)
}

type filterIntSliceTestCase struct {
	inSlice  []int
	outSlice []int
}

func prepareFilterIntSliceTestCases() []filterIntSliceTestCase {
	return []filterIntSliceTestCase{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{1, 2, 3, 5, 6, 7, 9}},
		{[]int{10, 12, 26, 32, 44}, []int{10, 12, 26, 32, 44}},
	}
}

func TestFilterIntSlice(t *testing.T) {
	testCases := prepareFilterIntSliceTestCases()
	for _, testCase := range testCases {
		inSlice := testCase.inSlice
		expectedSlice := testCase.outSlice
		filteredSlice := utils.FilterIntSlice(inSlice)
		for idx, value := range filteredSlice {
			if value != expectedSlice[idx] {
				t.Error(filteredSlice, "!=", expectedSlice)
			}
		}
	}
}

func TestPositiveBookSliceToMap(t *testing.T) {
	books := generateBookSlice()
	booksMap, _ := utils.BookSliceToMap(books)
	for _, book := range books {
		if _, ok := booksMap[book.UserId]; !ok {
			t.Error("Не найден UserId:", book.UserId)
		}
	}
}

func TestNegativeBookToSlice(t *testing.T) {
	testBook := book.Book{UserId: uint64(1), Title: "Title 1", Author: "Author 1", ISBN10: "1234567890", ISBN13: "1234567890123"}
	fmt.Println(book.String(&testBook))
	books := []book.Book{testBook, testBook}
	_, err := utils.BookSliceToMap(books)
	fmt.Println(err)
	if err == nil {
		t.Error("Не была обработана ситуациия с совпадающими UserId.")
	}
}
