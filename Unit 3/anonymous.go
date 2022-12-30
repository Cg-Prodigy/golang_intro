package main

import (
	"fmt"
	"math/rand"
	"time"
)

type colors func() uint8

func rgb() uint8{
	rand.Seed(time.Now().UnixNano())
	return uint8(rand.Intn(255))
}

func randomColor(value int, c colors) int{
	return int(c())+ value
}

func main(){
	color:=randomColor(10,rgb)
	fmt.Printf("%v \n",color)
}