package main

import (
	"fmt"
)

func main(){
	// planets := [...]string{
	// 	"Mercury",
	// 	"Venus",
	// 	"Earth",
	// 	"Mars",
	// 	"Jupiter",
	// 	"Saturn",
	// 	"Uranus",
	// 	"Neptune",
	// 	"Pluto",
	// }
	// terr:=planets[0:4]
	// gas:=planets[4:6]
	// ice:=planets[6:]
	// ice[0]="Asteroids"
	// fmt.Println(terr)
	// fmt.Println(gas)
	// fmt.Println(ice)
	// fmt.Println(planets)
	dwarfArray := [...]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dwarfSlice := dwarfArray[:]
	fmt.Printf("%T, %T",dwarfArray,dwarfSlice)
}