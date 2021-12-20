package top

// 4. 寻找两个正序数组的中位数
// https://leetcode-cn.com/problems/median-of-two-sorted-arrays/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	th := newTopHead()
	for _, v := range nums1 {
		th.add(v)
	}
	for _, v := range nums2 {
		th.add(v)
	}
	ans := 0.0
	pre := 0.0
	length := len(nums1) + len(nums2)
	mid := length / 2
	for i := 0; i <= mid; i++ {
		pre = ans
		ans = float64(th.remove())
	}
	if length%2 == 0 {
		ans = (ans + pre) / 2
	}
	return ans
}

// 从0开始
// 0  1  2   2N+1 2n+2
// 5, 6  2
// 小堆
type topHead struct {
	Value []int
	Size  int
}

func newTopHead() *topHead {
	return &topHead{Value: make([]int, 0), Size: 0}
}

func (t *topHead) add(v int) {
	t.Value = append(t.Value, v)
	t.Size++
	for c := t.Size - 1; c >= 0; {
		father := (c+1)>>1 - 1
		if father >= 0 && t.Value[c] < t.Value[father] {
			t.Value[c], t.Value[father] = t.Value[father], t.Value[c]
			c = father
		} else {
			break
		}
	}
}

func (t *topHead) remove() int {
	if t.Size == 0 {
		return 0
	}
	result := t.Value[0]
	t.Value[0] = t.Value[t.Size-1]
	t.Size--
	t.Value = t.Value[:t.Size]

	for c := 0; c < t.Size; {
		cLeftSon := c<<1 + 1
		cRightSon := cLeftSon + 1
		if cRightSon < t.Size && t.Value[cRightSon] < t.Value[cLeftSon] {
			cLeftSon = cRightSon
		}
		if cLeftSon < t.Size && t.Value[c] > t.Value[cLeftSon] {
			t.Value[c], t.Value[cLeftSon] = t.Value[cLeftSon], t.Value[c]
			c = cLeftSon
		} else {
			break
		}
	}
	return result
}
