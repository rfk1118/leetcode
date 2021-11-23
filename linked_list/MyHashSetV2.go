package linked_list

// 706. 设计哈希映射
// https://leetcode-cn.com/problems/design-hashmap/
import "container/list"

const v2base = 769

type entry struct {
	key, value int
}

type MyHashMapV2 struct {
	data []list.List
}

func ConstructorV2() MyHashMapV2 {
	return MyHashMapV2{make([]list.List, base)}
}

func (m *MyHashMapV2) hash(key int) int {
	return key % v2base
}

func (m *MyHashMapV2) Put(key, value int) {
	h := m.hash(key)
	for e := m.data[h].Front(); e != nil; e = e.Next() {
		if et := e.Value.(entry); et.key == key {
			e.Value = entry{key, value}
			return
		}
	}
	m.data[h].PushBack(entry{key, value})
}

func (m *MyHashMapV2) Get(key int) int {
	h := m.hash(key)
	for e := m.data[h].Front(); e != nil; e = e.Next() {
		if et := e.Value.(entry); et.key == key {
			return et.value
		}
	}
	return -1
}

func (m *MyHashMapV2) Remove(key int) {
	h := m.hash(key)
	for e := m.data[h].Front(); e != nil; e = e.Next() {
		if e.Value.(entry).key == key {
			m.data[h].Remove(e)
		}
	}
}
