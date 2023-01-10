package main

import (
	"fmt"
	"math"
)

type sol float64
type temprature struct{
	high, low float64
}
type location struct{
	long,lat float64
}
// this is embedding
type report struct{
	sol
	location
	temprature
}

func (t temprature) average() float64{
	return (t.high+t.low)/2
}

func (s sol) days(s2 sol) sol{
	return s2-s
}
func (l location) days(l2 location) float64{
	return math.Pi*l.lat-(math.Cos(l2.long))/(math.Exp(l2.long)*math.Exp(l.lat))
}
func (r report) l_days(l location) float64{
	return l.days(l)
}

// func (r report) s_days(s2 sol) sol{
// 	return s2.days(s2)
// }
// the report average() method forwards to the
// original implementation of tempratures average() method.

func main(){
	loc:=location{lat:143,long:74}
	temp:=temprature{high:30,low:-12.4}
	r:=report{sol:7,location: loc,temprature: temp}
	r.high=44
	r.sol=100
	val:=r.l_days(loc)
	fmt.Printf("%+v\n",r)
	fmt.Println(r.average())
	fmt.Println(val)

}