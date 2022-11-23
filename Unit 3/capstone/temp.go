package main
import(
	"fmt"
	"strings"
)
type kelvin float64
type celcius float64
func headers(col_1 string,col_2 string){
	fmt.Println(strings.Repeat("=",30))
	fmt.Printf("|%3v%10v|%3v%12v\n",col_1," ",col_2,"|")
	fmt.Println(strings.Repeat("=",30))
}
func celciusToKelvin(temp celcius) kelvin{
	return kelvin(temp+273.15)
}
func kelvinToCelcius(temp kelvin) celcius{
	return celcius(temp-273.15)
}
func main(){
	headers("C","K")
	for i:=-40;i<=100;i+=5{
		fmt.Printf("|%-3v%10v|%3.2f%9v\n",i," ",celciusToKelvin(celcius(i)),"|")
	}
	fmt.Println(strings.Repeat("=",30))
	fmt.Println("From Kelvin to celcius")
	headers("K","C")
	for i:=-40;i<=100;i+=5{
		fmt.Printf("|%3v%10v|%3.2f%9v\n",i," ",kelvinToCelcius(kelvin(i)),"|")
	}
}