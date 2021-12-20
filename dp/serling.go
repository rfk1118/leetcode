package dp

import (
	"math"

	"renfakai.com/leetcode/util"
)

var p []int = []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}

func cutRod(p []int, n int) int {
	if n == 0 {
		return 0
	}
	ans := math.MinInt64
	for i := 1; i <= n; i++ {
		ans = util.Max(p[n], p[i]+cutRod(p, n-i))
	}
	return ans
}

// 使用备忘录进行处理
func cutRodMemoized(p []int, n int) int {
	m := make([]int, n)
	for i := range m {
		m[i] = math.MinInt64
	}
	if n == 0 {
		return 0
	}
	var cutRodMemoizedAux func(p []int, aux []int, n int) int
	cutRodMemoizedAux = func(p []int, aux []int, n int) int {
		if aux[n] >= 0 {
			return aux[n]
		}
		if n == 0 {
			return 0
		}
		ans := math.MinInt32
		for i := 0; i <= n; i++ {
			ans = util.Max(p[n], p[i]+cutRodMemoizedAux(p, aux, n-i))
		}
		aux[n] = ans
		return ans
	}
	return cutRodMemoizedAux(p, m, n)
}

// 果然牛逼
func bottomUp(p []int, n int) int {
	var r []int = make([]int, n+1)
	r[0] = 0
	for j := 1; j <= n; j++ {
		q := math.MinInt64
		for i := 1; i <= j; i++ {
			// p[1] + r[0] = r[1]
			// p[1] + r[1]
			// p[2] + r[0] = r[2]
			// p[]
			q = util.Max(q, p[i]+r[j-i])
		}
		r[j] = q
	}
	return r[n]
}
