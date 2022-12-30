package main

import "fmt"


var cordinates struct{
	lat float64
	long float64
	alt int
}

type new_cordinates struct{
	lat float64
	long float64
	altitude int
}

var new_one new_cordinates
var new_two new_cordinates


func main(){
	new_one.lat=14.34566
	new_one.long=45.6754
	new_one.altitude=-5500

	new_two.lat=30.242
	new_two.long=-34.32
	new_two.altitude=3300

	fmt.Println(new_one,new_two)
}