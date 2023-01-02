package main

import "fmt"

var sType *string // the sType variable can be used with any variable of type string
func main(){
	loc:="Canada"
	sType:=loc
	fmt.Println(sType)
	fmt.Println(&sType)
}