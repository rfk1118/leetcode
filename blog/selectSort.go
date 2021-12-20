package blog

// selectSort
func selectSort(num []int) []int {
	for i := 0; i < len(num); i++ {
		minIndex := i
		for j := i + 1; j < len(num); j++ {
			if num[j] > num[minIndex] {
				minIndex = j
			}
		}
		num[i], num[minIndex] = num[minIndex], num[i]
	}
	return num
}
