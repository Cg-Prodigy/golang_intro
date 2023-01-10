package main
import "fmt"

// kelvin to Celsius

func kelvinToCelsius(val float64) float64{
	return val-273.15
}
func celsiusToFahrenheit(val float64) float64{
	return (val*9.0/5.0)+32.0
}
func kelvinToFahrenheit(val float64) float64 {
	return celsiusToFahrenheit(kelvinToCelsius(val))
}
func main(){
	kelvin:=145.23
	celsius:=kelvinToCelsius(kelvin)
	fmt.Println("Result: ",celsius)
}