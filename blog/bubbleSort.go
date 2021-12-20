package blog

// 冒泡排序
// https://www.renfakai.com/algorithms/sort/bubble-sort.html
func bubbleSort(num []int) []int {
	length := len(num)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if num[j] > num[j+1] {
				num[j], num[j+1] = num[j+1], num[j]
			}
		}
	}
	return num
}
