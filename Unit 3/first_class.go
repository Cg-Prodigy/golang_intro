package main

import (
	"fmt"
	"time"
	"math/rand"
)
type rgb uint8
type kelvin float64
func measureTemp(sample int, sensor func() kelvin){
	for i:=0; i<sample;i++{
		k:=sensor()
		fmt.Printf("%v K\n",k)
		time.Sleep(time.Second)
	}
	fmt.Println("Done!")
}
func fakeSensor() kelvin{
	return kelvin(rand.Intn(152)+150)
}
func main(){
	measureTemp(10,fakeSensor)
}