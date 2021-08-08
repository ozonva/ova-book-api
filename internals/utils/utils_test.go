package utils

import (
	"testing"
)

type testBatchedSliceCases struct {
	inSlice        []int
	batchSize      int
	expectedSlices [][]int
}

func prepareMakeBatchedSliceData() []testBatchedSliceCases {
	return []testBatchedSliceCases{
		{[]int{1, 2, 3, 4, 5}, 2, [][]int{{1, 2}, {3, 4}, {5}}},
		{[]int{1, 2, 3, 4, 5}, 3, [][]int{{1, 2, 3}, {4, 5}}},
		{[]int{1, 2, 3, 4, 5, 6}, 3, [][]int{{1, 2, 3}, {4, 5, 6}}},
	}
}

func TestMakeBatchedSlice(t *testing.T) {
	var testCases []testBatchedSliceCases = prepareMakeBatchedSliceData()

	for _, testCase := range testCases {
		expected := testCase.expectedSlices
		inSlice := testCase.inSlice
		batchSize := testCase.batchSize

		outSlice := MakeBatchedSlice(inSlice, batchSize)

		if len(expected) != len(outSlice) {
			t.Error("Не совпадает длина слайса, ожидается: ", len(expected), "Получено: ", len(outSlice))
		} else {
			for idx, value := range expected {
				for innerIdx, innerValue := range value {
					if innerValue != outSlice[idx][innerIdx] {
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
		resultMap := ReverseMap(caseMap)
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
	ReverseMap(inMap)
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
		filteredSlice := FilterIntSlice(inSlice)
		for idx, value := range filteredSlice {
			if value != expectedSlice[idx] {
				t.Error(filteredSlice, "!=", expectedSlice)
			}
		}
	}
}
