package main

import (
	"fmt"
	"data/graph"
)


func main() {
	n1:=1
	n2:=2
	n3:=3
	n4 :=4
	n5:=5
	 g :=graph.Graph{}
	g.New()
	g.Addedge(n1,n2)
	g.Addedge(n1,n3)
	g.Addedge(n1,n4)
	g.Addedge(n1,n5)
	g.Addedge(n2,n1)
	g.Addedge(n2,n3)
	g.Addedge(n2,n5)
	g.Addedge(n3,n1)
	g.Addedge(n3,n2)
	g.Addedge(n3,n5)
	g.Addedge(n3,n4)
	g.Addedge(n4,n1)
	g.Addedge(n4,n3)
	g.Addedge(n5,n1)
	g.Addedge(n5,n2)
	g.Addedge(n5,n3)
	no :=g.Cbfs()
	fmt.Println(no)
	// g.FindMnode()
	fmt.Println(g)
	
	
}

