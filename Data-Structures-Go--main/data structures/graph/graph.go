package graph

import "fmt"

// Graph map of graph class represent
type Graph struct{
	
	con map[int][]int
}

//New allways used first
func (g*Graph)New(){
	
	g.con = make(map[int][]int)
}


//Addedge function add a edge between two give nodw
func (g *Graph)Addedge(u,v int){
	g.con[u]=append(g.con[u],v)
}

//Bfs is workiing as expexteed
func (g *Graph)Bfs(){
	list := make(map[int]bool)
	for index :=range g.con{
		list[index]=false
	}
	for index := range g.con{
		for _,val :=range g.con[index]{
			if list[val] ==false{
				fmt.Println(val)
				list[val]=true
			}else{
				continue
			}
		}
	}
}
func (g *Graph)Cbfs()int{
	count :=0
	list := make(map[int]bool)
	for index :=range g.con{
		list[index]=false
	}
	for index := range g.con{
		for _,val :=range g.con[index]{
			if list[val] ==false{
				list[val]=true
				count++
			}else{
				continue
			}
		}
	}
	return count
}

//Dfs lets go down not working i dong kow how to solve it
func(g *Graph)Dfs(){
	list :=make(map[int]bool)
	for index :=range g.con{
		list[index]=false
	}
	for index:=range g.con{
		val:= g.con[index][0]
		for in:=range list{
			if list[val]==false{
				fmt.Println(val)
				list[val] = true
			}
			if list[val] ==true{
				continue
			}
			if list[in]==false{
				g.con[index]= g.con[index][1:]
				g.Dfs()
			}
		}
			
	}
}

func sum(ar []int)int{
	total:=0
	for _,val:=range ar{
		total +=val
	}
	return total
}

func filter(ar[]int)[]int{
	list := make(map[int]bool)
	z:=make([]int,0)
	for _,val := range ar{
		list[val]=false
	}
	for i:=0;i<len(ar);i++{
		if list[ar[i]]==false{
			z = append(z,ar[i])
			list[ar[i]]=true
		}else{
			continue
		}
	}
	return z
}


/*FindMnode is  a fuction to find mother node in a graph 
	mother node is a node through which all node can be reached
*/
func (g *Graph)FindMnode()int{
	list := make(map[int][]int)
	res := make(map[int]int)
	total :=0
	for index := range g.con{
		total+=index
		for _,val :=range g.con[index]{
			list[index] = append(list[index],g.con[index]...)
			list[index] = append(list[index],g.con[val]...)
		}
		list[index] =filter(list[index])		
	}
	for index := range list{
		res[index]= sum(list[index])
		if res[index] ==total{
			return index 
		}
	}
	return 0
} 

func remove(ar[]int,i int)[]int{
	z:= make([]int,0)
	z= append(z,ar[:i]...)
	z= append(z,ar[i+1:]...)
	return z
}
func (g*Graph)Kcores(k int){
	for index:= range g.con{
		if len(g.con[index])<k{
			delete(g.con,index)
			for in := range g.con{
				 for i,val:=range g.con[in]{
					 if val ==index{
						 g.con[in]= remove(g.con[in],i)
					 }
				 }
			}
		}else{
			continue
		}
	}
	fmt.Println(g.con)
}