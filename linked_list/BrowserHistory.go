package linked_list

// 1472. 设计浏览器历史记录
// https://leetcode-cn.com/problems/design-browser-history/
type BrowserHistory struct {
	homepage string
	V        []string
	currentV int
}

func BrowserHistoryConstructor(homepage string) BrowserHistory {
	r := BrowserHistory{
		homepage: homepage,
		V:        make([]string, 0),
		currentV: 0,
	}
	return r
}

func (bf *BrowserHistory) Visit(url string) {
	bf.V = bf.V[:bf.currentV]
	bf.V = append(bf.V, url)
	bf.currentV++
}

func (bf *BrowserHistory) Back(steps int) string {
	if bf.currentV < steps {
		bf.currentV = 0
		return bf.homepage
	}
	bf.currentV = bf.currentV - steps
	return bf.V[bf.currentV-1]
}

func (bf *BrowserHistory) Forward(steps int) string {
	i := len(bf.V)
	if i == 0 {
		return bf.homepage
	}
	if i-bf.currentV <= steps {
		bf.currentV = i
	} else {
		bf.currentV = bf.currentV + steps
	}
	return bf.V[bf.currentV-1]
}
