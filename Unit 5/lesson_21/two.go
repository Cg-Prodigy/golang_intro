package main
import "fmt"
type location struct{
	name,loc string
	color int8
	long,lat float64
}

var loc_two=location{}
func main(){
	loc_one:=location{lat:30.4,name:"Kenya",loc:"Central"}
	fmt.Printf("%+v",loc_one)
	fmt.Println(loc_two)
}