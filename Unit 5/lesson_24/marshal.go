package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type coordinate struct{
	d,m,s float64
	h rune
}
type location struct{
	Name string `json:"name"`
	Lat coordinate `json:"latitude"`
	Long coordinate `json:"longitude"`
}
func (c coordinate) String()string{
	return fmt.Sprintf("%v %v %.1f %c",c.d,c.m,c.s,c.h)
}
func (c coordinate) decimal() float64{
	sign:=1.0
	switch c.h{
	case 'S','s','W','w':
		sign*=-1
	}
	return sign*(c.d+c.m/60+c.s/3600)
}
func (c coordinate) MarshallJSON()([]byte,error){
	return json.Marshal(struct{
		DD float64 `json:"decimal"`
		DMS string `json:"dms"`
		D float64 `json:"degrees"`
		M float64 `json:"minutes"`
		S float64 `json:"seconds"`
		H string `json:"hemisphere"`

	}{
		DD: c.decimal(),
		DMS: c.String(),
		D: c.d,
		M: c.m,
		S: c.s,
		H: string(c.h)})
}

func main(){
	elysium:=location{
		Name:"Elysium Planitia",
		Lat:coordinate{d:4,m:30,s:0.0,h:'N'},
		Long:coordinate{d:135,m:54,s:0.0,h:'E'},
	}
	bytes,err:=json.MarshalIndent(elysium,""," ")
	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(bytes))
}