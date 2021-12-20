package sort

import "strconv"

// 506. 相对名次
// https://leetcode-cn.com/problems/relative-ranks/
func findRelativeRanks(score []int) []string {
	type Relative struct {
		Score, Index int
	}
	var tmp []*Relative
	for i, v := range score {
		tmp = append(tmp, &Relative{Score: v, Index: i})
	}
	var sort func(n []*Relative)
	sort = func(n []*Relative) {
		for i := 0; i < len(n); i++ {
			max := i
			for j := i + 1; j < len(n); j++ {
				if n[j].Score > n[max].Score {
					max = j
				}
			}
			n[i], n[max] = n[max], n[i]
		}
	}
	sort(tmp)
	var ans []string = make([]string, len(score))
	for i, v := range tmp {
		if i == 0 {
			ans[v.Index] = "Gold Medal"
		} else if i == 1 {
			ans[v.Index] = "Silver Medal"
		} else if i == 2 {
			ans[v.Index] = "Bronze Medal"
		} else {
			ans[v.Index] = strconv.Itoa(i + 1)
		}
	}
	return ans
}
