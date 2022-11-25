package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)
func modify_arr(arr [9]string){
	for i,elem:=range arr{
		fmt.Println(i,elem)
	}
}
func display(grid[9][9]uint8){
	for _,row:=range grid{
		fmt.Printf("%v\n",strings.Repeat("  - ",9))
		for _,col:=range row{
			fmt.Printf("| %v ",col)
		}
		fmt.Printf("|")
		fmt.Println()
		fmt.Printf("%v\n",strings.Repeat("  - ",9))
	}
}
func main(){
	rand.Seed(time.Now().UnixNano())
	var grid[9][9]uint8
	for i:=range grid{
		for j:=range grid[i]{
			grid[i][j]=uint8(rand.Intn(9)+1)
		}
	}
	display((grid))
}