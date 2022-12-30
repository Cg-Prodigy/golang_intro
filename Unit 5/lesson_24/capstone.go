package main

import (
	"math/rand"
)

type dog struct{
	name string
}
type Stringer interface{
	String() string
}
type animal interface{
	move() string
	eat() string
}
func (d dog) String() string{
	return d.name
}
func (d dog) move() string{
	switch rand.Intn(3){
	case 0:
		return "Runs"
	default:
		return "Sleeps"
	}
}
func (d dog) eat() string{
	switch rand.Intn(2){
	case 0:
		return "Meat"
	default:
		return "Dog Pellets"
	}
}

