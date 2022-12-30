package main

import (
	"math"
	"fmt"
)
type world struct{
	rad float64
}
type location struct{
	lat,long float64
}
type cordinates struct{
	d,m,s float64
	h rune
}
// constuctor for new location

func newLocation(lat, long cordinates) location{
	return location{lat.toLocation(),long.toLocation()}
}
func (w world) distance(p1,p2 location) float64{
	s1,c1:=math.Sincos(rad(p1.lat))
	s2,c2:=math.Sincos(rad(p2.lat))
	clong:=math.Cos(rad(p1.long-p2.long))
	return w.rad*math.Acos(s1*s2+c1*c2*clong)
}

func rad(cord float64) float64{
	return cord*math.Pi/180
}
func (cords cordinates) toLocation() float64{
	sign:=1.0
	switch cords.h{
	case 'W','w','s','S':
		sign*=-1
	}
	return sign* (cords.d+cords.m/60+cords.s/3600)
}
func main(){
	spirit:=[]cordinates{
		{d:14,m:34,s:6.2,h:'S'},
		{d:175,m:28,s:21.5,h:'E'},
	}
	curiosity:=[]cordinates{
		{d:4,m:35,s:22.2,h:'S'},
		{d:137,m:26,s:30.1,h:'E'},
	}
	mount_sharp:=[]cordinates{
		{d:5,m:4,s:48,h:'S'},
		{d:137,m:51,s:0,h:'E'},
	}
	olympus:=[]cordinates{
		{d:18,m:39,s:0,h:'N'},
		{d:226,m:12,s:0,h:'E'},
	}
	mecury:=world{rad:2439.7}
	earth:=world{rad:6371.0}
	mars:=world{rad:3389.5}
	s_location:=newLocation(spirit[0],spirit[1])
	c_location:=newLocation(curiosity[0],curiosity[1])
	o_location:=newLocation(olympus[0],olympus[1])
	m_location:=newLocation(mount_sharp[0],mount_sharp[1])
	distance_1:=mecury.distance(s_location,c_location)
	distance_2:=earth.distance(s_location,c_location)
	distance_3:=mars.distance(o_location,m_location)
	fmt.Println(distance_1)
	fmt.Println(distance_2)
	fmt.Println(distance_3)
}