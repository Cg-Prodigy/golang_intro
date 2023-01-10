package main

import (
	"fmt"
	"math"
)
type rover struct{
	gps
}
type gps struct{
	c_location,destination location
	world
}
type location struct{
	lat,long float64
	name string
}
type world struct{
	rad float64
}
func (l location) description()string{
	return fmt.Sprintf("Name: %v, longitude: %v, latitude: %v",l.name,l.long,l.lat)
}

func (w world) distance(p1,p2 location) float64{
	s1,c1:=math.Sincos(rad(p1.lat))
	s2,c2:=math.Sincos(rad(p2.lat))
	clong:=math.Cos(rad(p1.long-p2.long))
	return w.rad*math.Acos(s1*s2+c1*c2*clong)
}
func (loc gps) distance() float64{
	f_dest:=loc.world.distance(loc.c_location,loc.destination)
	return f_dest
}
func rad(c float64) float64{
	return c*math.Pi/180
}
func (loc gps) message() string{
	return fmt.Sprintf("%.1f km to %v",loc.distance(),loc.destination.description())
}

func main(){
	mars:=world{rad:3389.5}
	c_loc:=location{name:"Bradbury Landing",lat:-4.5895,long:137.4417}
	f_loc:=location{name:"Elysium Planitia",lat:4.5,long:135.9}
	c_gps:=gps{
		world: mars,
		c_location:c_loc,
		destination: f_loc,
	}
	r:=rover{gps:c_gps}
	fmt.Println(r.message())
}