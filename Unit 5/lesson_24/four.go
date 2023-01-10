// interfaces as types
package main

import (
	"fmt"
	"strings"
)

type talker interface{
	talk() string
}

type location struct{
	long,lat float64
}
type cordinates struct{
	location
	d,m,s float64
}
type custom string

func (c custom) talk() string{
	return "A talk method for type custom"
}
func (l location) talk() string{
	return fmt.Sprintf("Longitude: %v and Latitude: %v",l.long,l.lat)
}
// any argument passed to this function must satisify the talker interface
func shout(t talker){
	res:=strings.ToUpper(t.talk())
	fmt.Println(res)
}

func main(){
	loc:=location{lat:137.4,long:-14.45}
	// cust:=custom("Working hard")
	// fmt.Println(loc.talk())
	// fmt.Println(cust.talk())
	// shout(loc)
	// shout((cust))
	cords:=cordinates{d:100,m:14,s:22.2,location: loc}
	shout(cords)
}