package utils

func MakeBatchedSlice(inSlice []int, batchSize int) [][]int {
	sliceSize := len(inSlice)
	batches := sliceSize / batchSize

	var resultSlice [][]int
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
