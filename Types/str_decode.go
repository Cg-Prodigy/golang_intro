package main
import (
	"fmt"
	"unicode/utf8"
)

func main(){
	st:="abcdefghijklmnopqrstuvwxyz"
	fmt.Println(len(st))
	fmt.Println(utf8.RuneCountInString(st))
	// c,size:=utf8.DecodeRuneInString(st)
	// fmt.Printf("%c %v\n",c,size)
	// fmt.Println("Using range keyword")
	// for i,c:=range st{
	// 	fmt.Printf("%v %c %[2]v\n",i,c)
	// }
}