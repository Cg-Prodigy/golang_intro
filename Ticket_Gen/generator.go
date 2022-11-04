package main
import (
	"fmt"
	"strings"
	"math/rand"
	"time"
)

func main(){
	rand.Seed(time.Now().UnixNano())
	max:=100
	min:=20
	start:="SpaceX"
	t_type:="One-Way"
	fmt.Printf("%-20v %-6v %-14v %-6v\n","Spaceline","Days","Trip type","Price")
	fmt.Println(strings.Repeat("=",50))
	for i:=0;i<10;i++{
		days:=rand.Intn(50-10)+10
		price:=rand.Intn(max-min)+min
		fmt.Printf("%-20v %-6v %-14v $ %-6v\n",start,days,t_type,price)
		if start=="SpaceX"{
			start="Virgin Galactic"
		}else if start=="Virgin Galactic"{
			start="Space Adventures"
		}else{
			start="SpaceX"
		}
		if t_type=="One-Way"{
			t_type="Round-trip"
		}else{
			t_type="One-Way"
		}
	}

}