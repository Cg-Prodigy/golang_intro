package main

import (
	"fmt"
	"math"
)

type world struct{
	rad float64
}
type location struct{
	lat,long float64
}
var mars=world{rad:3389.5}

func (w world) distance(p1,p2 location) float64{
	s1,c1:=math.Sincos(rad(p1.lat))
	s2,c2:=math.Sincos(rad(p2.lat))
	clong:=math.Cos(rad(p1.long-p2.long))
	return w.rad*math.Acos(s1*s2+c1*c2*clong)
}

func rad(deg float64) float64{
	return deg*math.Pi/180
}

func main(){
	loc:=location{-14.5684,175.472636}
	loc_two:=location{-1.9462,354.4734}
	dist:=mars.distance(loc,loc_two)
	fmt.Printf("%.2f km\n",dist)
}