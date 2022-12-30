package main
import (
	"fmt"
)


func main(){
	st:="L fdph, L vdz, L frqtxhuhg"
	for i:=0;i<len(st);i++{
		char:=st[i]
		if char>='a' && char<='z'{
			char-=3
		}else if char>='A' && char<='Z'{
			char-=3
		}
		if char<0{
			char+=26
		}
		fmt.Printf("%c",char)
	}
	fmt.Println()
}