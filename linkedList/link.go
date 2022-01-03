package linkedlist

import "fmt"

type Node struct {
	Val  int
	Next *Node
}
type Llist struct {
	Head   *Node
	Length int
}

func InitNode(n int) *Node {
	return &Node{
		Val:  n,
		Next: nil,
	}
}

func (ll *Llist) AddHead(n int) {
	node := InitNode(n)
	if ll.Head != nil {
		fmt.Println("already head is presenrt")
		return
	}
	ll.Head = node
	ll.Length++
}

func (ll *Llist) RemoveDuplicates() {
	occurrence := make(map[int]int)

	for head := ll.Head; head.Next != nil; head = head.Next {

		if occurrence[head.Val] == 1 {
			fmt.Println("ok")
			head = head.Next.Next
			return
		}
		occurrence[head.Val]++

	}
	for index, val := range occurrence {
		fmt.Printf("index: %+v, val:%+v\n", index, val)
	}
}

func (ll *Llist) Insert(n int) {
	node := InitNode(n)
	head := ll.Head
	if ll.Head == nil {
		fmt.Println("EMPTY lIST EXCEPTIONt")
		return
	}
	for head.Next != nil {
		head = head.Next
	}
	head.Next = node
	ll.Length++
}

func (ll *Llist) Search(n int) int {
	head := ll.Head
	res := 0
	if ll.Head == nil {
		fmt.Println("EMPTY lIST EXCEPTIONt")
		return res
	}
	c := 1
	for head.Next != nil {
		c++
		if head.Next.Val == n {
			return c
		}
		head = head.Next
	}
	return 0
}

func (ll *Llist) SortedInsert(n int) {
	node := InitNode(n)
	head := ll.Head
	if ll.Head == nil {
		fmt.Println("EMPTY lIST EXCEPTIONt")
		return
	}
	for head.Next != nil && head.Next.Val < n {
		head = head.Next
	}
	node.Next = head.Next
	head.Next = node
	ll.Length++
}

func (ll *Llist) DeleteVal(n int) {
	// node := InitNode(n)
	head := ll.Head
	if ll.Head == nil {
		fmt.Println("EMPTY lIST EXCEPTIONt")
		return
	}
	for head.Next != nil {
		if head.Next.Val == n {
			head.Next = head.Next.Next
			ll.Length--
		}
		head = head.Next
	}
}
func (ll *Llist) Reverse() {
	head := ll.Head
	prev, next := &Node{}, &Node{}
	for head != nil {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}
	ll.Head = prev

}
func (ll *Llist) Iterate() {
	head := ll.Head
	for val := head; val != nil; val = val.Next {
		fmt.Println(val.Val)
	}
}
