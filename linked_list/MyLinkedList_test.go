package linked_list

import (
	"fmt"
	"testing"
)

func TestMyLinkedList_Get(t *testing.T) {
	constructor := MyLinkedListConstructor()
	//constructor.AddAtHead(1)
	//constructor.AddAtTail(3)
	//constructor.AddAtIndex(1,2)
	//node := constructor.Get(1)
	//fmt.Println(node)
	//constructor.DeleteAtIndex(1)
	//node2 := constructor.Get(1)
	//fmt.Println(node2)
	//constructor.AddAtIndex(0,10)
	//constructor.AddAtIndex(0,20)
	//constructor.AddAtIndex(1,30)
	//get := constructor.Get(0)
	//fmt.Println(get)

	// ["MyLinkedList","addAtHead","get","addAtHead","addAtHead","deleteAtIndex","addAtHead","get","get","get","addAtHead","deleteAtIndex"]
	//[[],[4],[1],[1],[5],[3],[7],[3],[3],[3],[1],[4]]
	//constructor.AddAtHead(4)
	//get := constructor.Get(1)
	//fmt.Println(get)
	//constructor.AddAtHead(1)
	//constructor.AddAtHead(5)
	//constructor.DeleteAtIndex(3)
	//constructor.AddAtHead(7)
	//get3 := constructor.Get(3)
	//fmt.Println(get3)
	//get33 := constructor.Get(3)
	//fmt.Println(get33)
	//get333 := constructor.Get(3)
	//fmt.Println(get333)
	//constructor.AddAtHead(1)
	//constructor.DeleteAtIndex(4)
	//fmt.Println(constructor)

	//["MyLinkedList","addAtHead","addAtTail","addAtIndex","get","deleteAtIndex","get","get","deleteAtIndex","deleteAtIndex","get","deleteAtIndex","get"]
	//[[],[1],[3],[1,2],                                    [1],     [1],          [1],  [3],       [3],         [0],           [0],      [0],      [0]]
	constructor.AddAtHead(1)
	constructor.AddAtTail(3)
	constructor.AddAtIndex(1, 2)
	get := constructor.Get(1)
	fmt.Println(get)
	constructor.DeleteAtIndex(1)
	get1 := constructor.Get(1)
	ge33 := constructor.Get(3)
	fmt.Println(get1)
	fmt.Println(ge33)
	constructor.DeleteAtIndex(3)
	constructor.DeleteAtIndex(0)
	get0 := constructor.Get(0)
	fmt.Println(get0)
	constructor.DeleteAtIndex(0)
	get00 := constructor.Get(0)
	fmt.Println(get00)

}
