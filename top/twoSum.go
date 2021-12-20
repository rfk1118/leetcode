package top

// 数组并没有说有序
func twoSum(nums []int, target int) []int {
	type V struct {
		Value, Index int
	}

	var tmp []*V
	for i, v := range nums {
		tmp = append(tmp, &V{Value: v, Index: i})
	}
	// 我就得手写排序，我就要练习选择
	for i := 0; i < len(tmp); i++ {
		min := i
		for j := i + 1; j < len(tmp); j++ {
			if tmp[j].Value < tmp[min].Value {
				min = j
			}
		}
		tmp[i], tmp[min] = tmp[min], tmp[i]
	}
	var ans []int = make([]int, 0)
	for i, j := 0, len(tmp)-1; i < j; {
		if tmp[i].Value+tmp[j].Value == target {
			ans = append(ans, tmp[i].Index)
			ans = append(ans, tmp[j].Index)
			break
		} else if tmp[i].Value+tmp[j].Value > target {
			j--
		} else if tmp[i].Value+tmp[j].Value < target {
			i++
		}
	}
	return ans
}
