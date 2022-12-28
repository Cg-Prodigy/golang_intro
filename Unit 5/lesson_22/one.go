package main

import "fmt"

type cordinates struct{
	d,m,s float64
	h rune
}
type location struct{
	lat,long float64
}
// constuctor function
func newLocation(lat,long cordinates) location{
	return location{lat.to_decimal(),long.to_decimal()}
}
func (c cordinates) to_decimal() float64{
	sign:=1.0
	switch c.h{
	case 'W','w','S','s':
		sign*=-1
	}
	return sign *(c.d+c.m/60+c.s/3600)
}


func main(){
	b_landing:=cordinates{d:4,m:35,s:22.2,h:'S'}
	a_landing:=cordinates{d:137,m:26,s:30.12,h:'E'}
	cords_1:=newLocation(b_landing,a_landing)
	fmt.Println(cords_1)

}