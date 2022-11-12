package main
import "fmt"

type fahrenheit float64
type celsius float64
type kelvin float64
type sound string
func (val fahrenheit) celsius() celsius{
	return celsius((val-32)*5.0/9.0)
}
func (val kelvin) celsius() celsius{
	return celsius(val-273.15)
}
func (val celsius) kelvin() kelvin{
	return kelvin(val+273.15)
}
func (s sound) sounds(age int,legs int) sound{
	if age<=4 && legs==2{
		return sound("Ostrich")
	}else{
		return sound("Cow")
	}
}
func main()  {
	var temp celsius=50.45
	var temp_2 fahrenheit=30
	var t_sound sound="Feather"
	fmt.Println(temp.kelvin())
	fmt.Println(temp_2.celsius().kelvin())
	fmt.Println(t_sound.sounds(5,1))
}