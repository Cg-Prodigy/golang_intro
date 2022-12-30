package main

import (
	"fmt"
	"math"
)

var t interface{
	talker() string
}

type location struct{
	lat,long float64
}
type cordinates struct{
	d,m,s float64
	h rune
}

func (l location) talker() string{
	return "This is a talker for location type."
}
func (l location) crazyMath() float64{
	return math.Cos(l.lat)+math.Sin(l.long)
}
func (c cordinates) talker() string{
	return "This is a talker for cordinate type."
}
func newLocation(cord1,cord2 cordinates) location{
	return location{cord1.radians(),cord2.radians()}
}
func (c cordinates) radians()float64{
	sign:=1.0
	switch c.h{
	case 'S','s','W','w':
		sign*=-1
	}
	return sign*(c.d+c.m/60*c.s/3600)
}


func main(){
	cord_1:=cordinates{d:100,m:60,s:22.4,h:'S'}
	t=cord_1
	fmt.Println(t.talker())
	cord_2:=cordinates{d:137,m:120,s:43,h:'W'}
	loc:=newLocation(cord_1,cord_2)
	fmt.Println(loc.crazyMath())

}