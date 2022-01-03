package linkedlist

//sorting(insertion sort,selection),summing,counting(how many nodes),swapping(accepts two input),
//Sortin is an algorithm for sorting an linked list usiing insertion sortworkig need improvement
func (l *Llist) Sortin() {
	//this algorithm is not working
	var node *Node
	var node1 *Node
	for node = l.Head; node != nil; node = node.Next {
		for node1 = l.Head.Next; node1 != nil; node1 = node1.Next {
			if node1 !=nil&&node!=nil{
				if node1.Val < node.Val {
					swap(&node.Val, &node.Next.Val)

				}
			}
			// if node.Val<node.Next.Val{
			// 	break
			// }
			// continue

		}
	}
}

//Sortsel selection sort algorithm for Llist
func (l *Llist) Sortsel() {

}

//Sum sums up all node Val off
func (l *Llist) Sum() int {
	total := 0
	var node *Node
	for node = l.Head; node != nil; node = node.Next {
		if node != nil {
			total += node.Val
		}
	}

	return total
}

//util function for swapping algorithm
func swap(a *int, b *int) {
	var tmp = *a
	*a = *b
	*b = tmp
}

//Swap algo swaps two gven nde values
func (l *Llist) Swap(val1 int, val2 int) {
	var node *Node
	var node1 = &Node{}
	var node2 = &Node{}
	for node = l.Head; node != nil; node = node.Next {
		if node.Val == val1 {
			node1 = node
		}
		if node.Val == val2 {
			node2 = node
		}
		var point int
		point = node1.Val
		//swapping works but we lose 2nd variable
		swap(&node1.Val, &node2.Val)
		node2.Val = point
	}
}

//Count is working
func (l *Llist) Count() int {
	var node *Node
	count := 0
	for node = l.Head; node != nil; node = node.Next {
		if node != nil {
			count++
		} else {
			return 0
		}
	}
	return count
}
