package main

import "fmt"

var admin *string

func main(){
	scholes:="Christopher J. Scholes"
	admin=&scholes
	start:=admin

	
	bolden:="Charles F. Bolden"
	admin=&bolden
	*admin="Charles Frank Bolden Jr."
	major:=admin
	*major="Major General Charles Frank Bolden Jr."
	lightfoot:="Robert M. Lightfoot Jr."
	admin=&lightfoot
	// *major="Ego Brian"
	fmt.Println(&bolden,bolden)
	fmt.Println(major,*major)
	fmt.Println(admin,*admin)
	fmt.Println(start,*start)

}