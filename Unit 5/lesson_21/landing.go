package main

import (
	"encoding/json"
	"fmt"
)

type location struct {
	Name string `json:"landing Site"`
	Lat float64 `json:"latitude"`
	Long float64 `json:"longitude"`
}

func main(){
	locations := []location{
		{Name: "Bradbury Landing", Lat: -4.5895, Long: 137.4417},
		{Name: "Columbia Memorial Station", Lat: -14.5684, Long: 175.472636},
		{Name: "Challenger Memorial Station", Lat: -1.9462, Long: 354.4734},
	}
	bytes,err:=json.MarshalIndent(locations,"  ","")
	if err==nil{
		fmt.Println(string(bytes))
	}
}