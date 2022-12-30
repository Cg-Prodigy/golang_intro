// interface as a variable

package main

import "fmt"

var t interface{
	talk() string
}

type one struct{
	lat,long float64
}
type custom int

func (c custom) talk() string{
	return "Type custom talking"
}
func (o one) talk() string{
	return fmt.Sprintf("Longitude: %v and Latitude: %v",o.long,o.lat)
}


func main(){
	t=custom(14)
	fmt.Println(t.talk())
	t=one{long:14.5,lat:-13.54}
	fmt.Println(t.talk())
}