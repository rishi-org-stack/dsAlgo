package main

import (
	cs "dsAlgo/problems/cses"
	"fmt"
)

func main() {
	// ll := &llist.Llist{}
	// ll.AddHead(1)
	// ll.Insert(2)
	// ll.Insert(5)
	// ll.Insert(4)
	// ll.Insert(4)
	// ll.Insert(3)
	// ll.Insert(4)
	// fmt.Println("before reversed :")
	// ll.Iterate()
	// ll.Reverse()
	// fmt.Println("after reversed")
	// ll.Iterate()
	c := cs.Playlist([]int{1, 2, 1, 3, 2, 7, 4, 2})
	// arr := []int{30, 60, 75}
	// ar := cs.Pop(arr, 0)
	fmt.Println(c)
}
