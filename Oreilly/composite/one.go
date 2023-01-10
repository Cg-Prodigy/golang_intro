package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// the zero value of a map is nil
// that means you cannot assing to a nil map
var nilMap map[string]int //this is a nil map
// unlike a nil map, a map created usuing an map literal can be assigned new values
// maps are not comparable. You can only check wether a map is nil as in, map==nil
// However you cannot check wether two maps contain identical keys and values.
// You cannot use slices and maps as keys to maps.
func populateMap(m map[string]int){
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<100;i++{
		val:=string(strconv.Itoa(rand.Intn(100)))
		m[val]=rand.Intn(50)*20
	}
}
func popTwo(lst []string) map[string]int{
	rand.Seed(time.Now().UnixNano())
	newMap:=make(map[string]int,len(lst))
	for _,value:= range lst{
		val:=rand.Intn(100)
		newMap[value]=val
	}
	return newMap
}

func main(){
	lst:=[]string{"One","Two","Three","Four"}
	newMap:=popTwo(lst)
	fmt.Println(newMap)
	// val,ok:=newMap["Four"]
	// fmt.Println(val,ok)
	delete(newMap,"Two")

}

// To delete elements from a map, the delete function is used
// The delete function takes the map and key as arguments
// If the key is not present, nothing happens, else the key-value pair is deleted
// delete function returns nothing.