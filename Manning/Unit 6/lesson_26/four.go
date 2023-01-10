package main

import "fmt"

type person struct{
	name,supepower string
	age int
}

func main(){
	ego:=&person{
		name:"Ego",
		age:10,
	}
	ego.supepower="Intelligence"
	fmt.Printf("%+v\n",ego)
}