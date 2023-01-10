package main

import (
	"fmt"
	"math/rand"
	"time"
)
func randColor() []int{
	rand.Seed(time.Now().UnixNano())
	color_lst:=make([]int,0)
	for i:=0;i<3;i++{
		color:=rand.Intn(100)
		func(j int){
			if j==77{
				j=rand.Intn(100)
			}
		}(color)
		color_lst=append(color_lst, color)

	}
	return color_lst
}

func main(){
	color:=randColor()
	fmt.Println(color)
}