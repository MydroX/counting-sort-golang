package sorting

func Sort(numbers []int) []int {
	maxInt := getMaxIntArray(numbers)
	countArray := make([]int, maxInt+1)
	result := make([]int, len(numbers)+1)

	countMap := countInArray(numbers)
	for i := 1; i < maxInt+1; i++ {
		countArray[i] = countMap[i] + countArray[i-1]
	}

	for j := 0; j < len(numbers); j++ {
		result[countArray[numbers[j]]] = numbers[j]
		countArray[numbers[j]] = countArray[numbers[j]] - 1
	}

	return result[1:]
}

func countInArray(arr []int) map[int]int {
	countMap := make(map[int]int)
	for _, elt := range arr {
		countMap[elt] += 1
	}
	return countMap
}

func getMaxIntArray(arr []int) int {
	maxInt := arr[0]
	for _, elt := range arr {
		if maxInt < elt {
			maxInt = elt
		}
	}
	return maxInt
}
