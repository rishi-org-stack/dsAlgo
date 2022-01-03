package stack

import "fmt"

type Stack struct {
	num  int
	elem []interface{}
}

func (s *Stack) Init(n int) {
	s.num = n
	s.elem = make([]interface{}, 0)
}

func (s Stack) Size() int {
	return len(s.elem)
}
func (s Stack) Push(n interface{}) Stack {
	var z []interface{}
	var res Stack
	if s.Size()+1 <= s.num {
		z = append(z, n)
		z = append(z, s.elem...)
	} else {
		fmt.Println("stack overflow")
	}
	res.num=s.num
	res.elem=z
	return res
}

func (s Stack)Pop()interface{}{
	return s.elem[0]
}

func (s Stack)Isempty()bool{
	c:=0
	for i:=0;i<len(s.elem);i++{
		if s.elem[i]==nil{
			c++
		}
	}
	if c==len(s.elem){
		return true
	}
	return false
}