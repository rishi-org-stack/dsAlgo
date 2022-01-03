package queue

import "fmt"

type Queue struct {
	num  int
	elem []interface{}
}

func (q Queue) Init(n int) {
	q.num = n
	q.elem = make([]interface{}, 0)
}

func (q Queue) Enqueue(i interface{}) Queue {
	if len(q.elem)+1 <= q.num {
		q.elem = append(q.elem, i)
	} else {
		fmt.Println("No Space is left")
	}
	return q
}

func (q Queue)Dequeue()Queue{
	q.elem=  q.elem[:len(q.elem)-1]
	return q
}


func (q Queue)Isempty()bool{
	// c:=0
	if len(q.elem)==0{
		return true
	}
	return false
}