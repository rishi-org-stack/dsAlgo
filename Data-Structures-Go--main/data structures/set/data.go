package set

import "fmt"
//Algebric ,Union,Intersection, Deletion,Addelement,Contaielement,
type Set struct{
	mapping map[int]bool
}
func (s *Set) New(){
	s.mapping = make(map[int]bool)
}

func (s *Set) Contains(elem int)bool{
	var boolean bool
	boolean = s.mapping[elem]
	if boolean ==true {
		return boolean
	}
	boolean = false
	return boolean 
}

func (s *Set)Addelem(elem int){
	if !s.Contains(elem){
		s.mapping[elem] = true
	}else{
		fmt.Println("already exist")
	}
	
}

func (s *Set)DeleteElem(elem int){
	if s.Contains(elem){
		delete(s.mapping,elem)
	}else{
		fmt.Println("no such element exist")
	}
	
}

func (s*Set)Intersection(as *Set)*Set{
	as.New()
	var inser = &Set{}
	for in ,elem:=range s.mapping{
		if elem == as.mapping[in] && elem == true{
			inser.Addelem(in)
		}
	}
	return inser
}
func (s*Set)Union(as *Set) *Set{
	var res =&Set{}
	for in,val:=range s.mapping{
		if val == as.mapping[in]{
			res.Addelem(in)
		}
	}
	return res
}