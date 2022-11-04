package main
import (
	"fmt"
	"math/rand"
	"time"
)
func main(){
	rand.Seed(time.Now().UnixNano())
	month:=rand.Intn(12)+1
	dayInMonth:=31
	switch month{
	case 2:
		dayInMonth=28
	case 4,6,9,11:
		dayInMonth=30
	}
	day:=rand.Intn(dayInMonth)+1
	fmt.Println(month,day)

}