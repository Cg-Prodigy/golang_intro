package main

import (
	"fmt"
)

// anonymous functions
// functions defined inside other functions
// and assigned to new variable


func main(){
	for i:=0;i<10;i++{
		func (j int) bool{
			fmt.Printf("%v is actually i\n",j)
			return true
		}(i)
	}
}