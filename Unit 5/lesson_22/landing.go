package main

import "fmt"


type location struct{
	d,m,s float64
	h rune
}
type cords struct{
	lat,long float64
}

func newLocation(lat,long location) cords{
	return cords{lat.decimal(),long.decimal()}
}

func (c location) decimal() float64{
	sign:=1.0
	switch c.h{
	case 'W','w','s','S':
		sign*=-1
	}
	return sign*(c.d+c.m/60+c.s/3600)
}

func main(){
	spirit:=[]location{{d:14,m:34,s:6.2,h:'S'},{d:175,m:28,s:21.5,h:'E'}}
	fmt.Println(spirit[0].decimal())
}