package main
import (
	"fmt"
	"math"
)

func  main()  {
	v:=100
	if v>=0 && v<=math.MaxUint8{
		v:=uint8(v)
		fmt.Printf("%v %[1]T",v)
	}else{
		fmt.Println("Out of range")
	}
}
