package main
import(
	"fmt"
	"math/rand"
	"time"
)
func main(){
	rand.Seed(time.Now().UnixNano())
	var balance float32=0.0
	for {
		i:=rand.Intn(3)
		if i==0{
			balance+=0.05
		}else if i==1{
			balance+=0.1
		}else{
			balance+=0.25
		}
		if balance>=20{
			break
		}
		fmt.Printf("$%.2f\n",balance)
	}
}