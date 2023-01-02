package main

import "fmt"

func main(){
	answer:=43
	address:=&answer
	fmt.Println(address)
	fmt.Println(*address)
	fmt.Printf("%T\n",address)
	fmt.Printf("%T\n",*address)
}