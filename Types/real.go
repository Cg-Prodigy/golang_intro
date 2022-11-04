package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	piggBank:=0.0
	nickel:=0.05
	dimes:=0.10
	quarter:=0.25
	for {
		r_int:=rand.Intn(3)
		if r_int==0{
			piggBank+=nickel
		}else if r_int==1{
			piggBank+=dimes
		}else{
			piggBank+=quarter
		}
		if piggBank>=20{
			break
		}
		fmt.Printf("Piggy Bank Balance: %.2f\n",piggBank)
	}
}
