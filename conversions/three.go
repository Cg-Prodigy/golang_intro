package main
import (
	"fmt"
)

func main(){
	value:=false
	v_text:=fmt.Sprintf("%v",value)
	var v_two bool;
	v_two=(v_text=="true")
	fmt.Printf("%v %[1]T\n",value)
	fmt.Printf("%v %[1]T\n",v_text)
	fmt.Printf("%v %[1]T\n",v_two)
}