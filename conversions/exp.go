package main
import(
	"fmt"
)
func main()  {
	st:="true"
	var b bool;
	switch st {
	case "true","yes","1":
		b=true
	case "false","no","0":
		b=false
	}
	fmt.Println(b)
}