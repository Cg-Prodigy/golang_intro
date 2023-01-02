package main

import "fmt"

// import (
// 	"fmt"
// )

// func makeMult(base int) func(int) int{
// 	return func(i int) int {
// 		return base*i
// 	}
// }

// func main(){
// 	twoBase:=makeMult(2)
// 	threBase:=makeMult(3)
// 	for i:=0;i<5;i++{
// 		fmt.Println(twoBase(i),threBase(i))
// 	}
// }

func modMap(m map[int]string){
	m[2]="hello"
	m[3]="goodbye"
	delete(m,1)
}
func modSlice(s []int){
	for i,v:=range s{
		s[i]=v*2
	}
	s=append(s, 10)
}

func main(){
	m:=map[int]string{
		1:"first",
		2:"second",
	}
	modMap(m)
	fmt.Println(m)
	s:=[]int{1,2,3}
	modSlice(s)
	fmt.Println(s)
}