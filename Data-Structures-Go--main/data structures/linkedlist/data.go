package linkedlist

import (
	"fmt"
)

type Node struct {
	Val    int
	Next   *Node
	Before *Node
}
type Llist struct {
	Head *Node
}

func (ll *Llist) Addtohead(val int) {
	var nod Node
	nod.Val = val
	if nod.Next == nil {
		nod.Next = ll.Head
	}
	ll.Head = &nod
}
func (ll *Llist) Iterate() {
	var nod *Node
	for nod = ll.Head; nod != nil; nod = nod.Next {
		fmt.Println(nod.Val)
	}
}
func (ll *Llist) Lastnode() *Node {
	var nod *Node
	var lnod *Node
	for nod = ll.Head; nod != nil; nod = nod.Next {
		if nod.Next == nil {
			lnod = nod
		}
	}
	return lnod
}
func (ll *Llist) Addtoend(val int) {
	var nod = &Node{}
	nod.Val = val
	var last *Node
	last = ll.Lastnode()
	if last != nil {
		last.Next = nod
	}
	last = nod
}
func (ll *Llist) Find(val int) *Node {
	var node *Node
	for nod := ll.Head; nod != nil; nod = nod.Next {
		// count++
		if nod.Val == val {
			node = nod
			// break
		}
	}
	return node
}
func (ll *Llist) Addafter(val1 int, val2 int, val3 int) {
	var nod = &Node{} //*Node will not work
	nod.Val = val3
	for node := ll.Head; node != nil; node = nod.Next {
		if node.Next.Val == val1 && node.Before.Val == val2 {
			node.Next = nod
		}
	}

}
