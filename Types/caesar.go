package main
import(
	"fmt"
)

func main(){
	m:="uv vagreangvbany fcnpr fgngvba"
	for i:=0;i<len(m);i++{
		c:=m[i]
		if c>='a' && c<='z'{
			c+=13
			if c>'z'{
				c-=26
			}
		}
		fmt.Printf("%c",c)
	}
	fmt.Println()
}