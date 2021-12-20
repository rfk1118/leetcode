package linked_list

// https://leetcode-cn.com/problems/design-twitter/
// 355. 设计推特

import "container/heap"

// Twitter 这里代表推特整个系统
type Twitter struct {
	timestamp int
	userMap   map[int]*Userer
}

func TwitterConstructor() Twitter {
	return Twitter{0, make(map[int]*Userer, 0)}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	if _, ok := this.userMap[userId]; !ok {
		// 如果还没初始化的化进行初始化，这里不能return，因为后面要新增
		this.userMap[userId] = NewUser(userId)
	}
	u := this.userMap[userId]
	u.post(tweetId, this.timestamp)
	this.timestamp++
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	var (
		res []int
	)
	if _, ok := this.userMap[userId]; !ok {
		return res
	}
	hp := &TimeQueue{}
	users := this.userMap[userId].Followed
	for u, _ := range users {
		twt := this.userMap[u].Head
		if twt == nil {
			continue
		}
		heap.Push(hp, twt)
	}
	for hp.Len() != 0 {
		if len(res) == 10 {
			break
		}
		twt := heap.Pop(hp).(*Tweet)
		res = append(res, twt.id)
		if twt.nex != nil {
			heap.Push(hp, twt.nex)
		}
	}
	return res
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := this.userMap[followerId]; !ok {
		this.userMap[followerId] = NewUser(followerId)
	}
	if _, ok := this.userMap[followeeId]; !ok {
		this.userMap[followeeId] = NewUser(followeeId)
	}
	u := this.userMap[followerId]
	u.follow(followeeId)
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if _, ok := this.userMap[followerId]; ok {
		u := this.userMap[followerId]
		u.unfollow(followeeId)
	}
}

type Userer struct {
	id       int
	Followed map[int]flag
	Head     *Tweet
}

type flag struct{}

func (u *Userer) follow(id int) {
	u.Followed[id] = flag{}
}

func (u *Userer) unfollow(id int) {
	if id != u.id {
		delete(u.Followed, id)
	}
}

func (u *Userer) post(tweetId, timestamp int) {
	twt := NewTweet(tweetId, timestamp)
	twt.nex = u.Head
	u.Head = twt
}

func NewUser(id int) *Userer {
	temp := &Userer{
		id,
		make(map[int]flag),
		nil,
	}
	temp.follow(id)
	return temp
}

type Tweet struct {
	id   int
	time int
	nex  *Tweet
}

func NewTweet(id int, time int) *Tweet {
	return &Tweet{
		id,
		time,
		nil,
	}
}

type TimeQueue []*Tweet

func (hp TimeQueue) Len() int            { return len(hp) }
func (hp TimeQueue) Less(i, j int) bool  { return hp[i].time > hp[j].time }
func (hp TimeQueue) Swap(i, j int)       { hp[i], hp[j] = hp[j], hp[i] }
func (hp *TimeQueue) Push(x interface{}) { *hp = append(*hp, x.(*Tweet)) }
func (hp *TimeQueue) Pop() interface{} {
	old := *hp
	x := (*hp)[len(old)-1]
	*hp = old[:len(old)-1]
	return x
}
