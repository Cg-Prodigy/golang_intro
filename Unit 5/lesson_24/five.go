package main

import (
	"fmt"
	"time"
)
type sol int
func (s sol) YearDay() int{
	return int(s%356)
}
func (s sol) Hour() int{
	return int(s%24)
}
type stardater interface{
	YearDay() int
	Hour() int
}
func stardate(t stardater) float64{
	doy:=float64(t.YearDay())
	h:=float64(t.Hour())/24.0
	return 1000+doy+h
}

func main(){
	day:=time.Date(2020,1,3,17,12,0,0,time.UTC)
	s:=sol(1444)
	fmt.Println(stardate(s))
	fmt.Println(stardate(day))
}