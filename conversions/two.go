package main

import (
	"fmt"
	"strconv"
	// "strconv"
)
func main(){
	num:=100000
	st:="100000"
	res:=strconv.Itoa(num)
	res_two,err:=strconv.Atoi(st)
	if err!=nil{
		fmt.Println("Error occured during conversion")
	}else{
		fmt.Printf("%v %[1]T\n",res_two)
	}
	fmt.Printf("%v %[1]T\n",res)
	val:=fmt.Sprintf("%v",num)
	fmt.Println(val)
}