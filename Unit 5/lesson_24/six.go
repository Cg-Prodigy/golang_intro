package main

import "fmt"

type Stringer interface{
	String() string
}
type location struct{
	long,lat cordinate
}
type cordinate struct{
	d,m,s float64
	h rune
}

func (c cordinate) String() string{
	return fmt.Sprintf("is at %vº%v'%v'' %c",c.d,c.m,c.s,c.h)
}
func (l location) String() string{
	return fmt.Sprintf("Longitude: %v, Latitude: %v",l.long,l.lat)
}

func main(){
	loc:=location{
		lat:cordinate{4, 30, 0.0, 'N'},
		long:cordinate{135, 54, 0.0, 'E'},
	}
	fmt.Println(loc)
}