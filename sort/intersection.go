package sort

// 排序后进行处理
// 超出时间限制
func intersection(nums1 []int, nums2 []int) []int {
	var sort func(nums []int) []int
	sort = func(nums []int) []int {
		for i := 0; i < len(nums); i++ {
			for j := 0; j < len(nums)-1-i; j++ {
				if nums[j] > nums[j+1] {
					nums[j], nums[j+1] = nums[j+1], nums[j]
				}
			}
		}
		return nums
	}
	x := sort(nums1)
	y := sort(nums2)
	var ans []int
	for i, j := 0, 0; i < len(x) && j < len(y); {
		if x[i] > y[j] {
			j++
		} else if x[i] < y[j] {
			i++
		} else {
			for i < len(x)-1 {
				if x[i] == x[i+1] {
					i++
				}
			}
			for j < len(y)-1 {
				if y[j] == y[j+1] {
					j++
				}
			}
			ans = append(ans, x[i])
			i++
			j++
		}
	}
	return ans
}

func intersectionMap(nums1 []int, nums2 []int) []int {
	m := make(map[int]bool, 0)
	mAdd := make(map[int]bool, 0)
	for _, v := range nums1 {
		m[v] = true
	}
	var ans []int
	for _, v := range nums2 {
		if m[v] && !mAdd[v] {
			ans = append(ans, v)
			mAdd[v] = true
		}
	}
	return ans
}
