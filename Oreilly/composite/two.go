package main

import "fmt"

// sets using maps

func main(){
	set:=map[int]bool{}
	lst:=[]int{10,20,3,4,5,4,10,15,13,45,90,45,30}
	for _,val:=range lst{
		set[val]=true
	}
	fmt.Println(len(set),len(lst))
}