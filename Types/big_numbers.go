package main
import (
	"fmt"
)

// func main(){
// 	lt:=new(big.Int)
// 	st:="9.461"+strings.Repeat("0",15)
// 	lt.SetString(st,10)
// 	dist:=new(big.Int)
// 	st_t:="236"+strings.Repeat("0",18)
// 	dist.SetString(st_t,10)
// 	result:=new(big.Int)
// 	result.Div(dist,lt)
// 	fmt.Println("Result is: ",result)

// }
// func main(){
// 	const light_y=9.461e15
// 	const distance=2.4e19
// 	const result=distance/light_y
// 	fmt.Println("Result is: ", result)
// }
func main(){
	fmt.Println(2.4e19/9.461e15)
}