package main

import "fmt"

type temprature struct{
	high, low float64
}
type location struct{
	long,lat float64
}

type report struct{
	sol int
	loc location
	temp temprature
}

func (t temprature) average() float64{
	return (t.high+t.low)/2
}

// the report average() method forwards to the
// original implementation of tempratures average() method.
func (r report) average() float64 {
	return r.temp.average()
}


func main(){
	loc:=location{lat:143,long:74}
	temp:=temprature{high:30,low:-12.4}
	r:=report{sol:7,loc:loc,temp:temp}

	fmt.Printf("%+v\n",r)
	fmt.Println(r.average())
}