package main

import (
	"fmt"
	"sort"
)

// passing functions as parameters
type Person struct{
	FirstName, LastName string
	age int
}

func main(){
	people:=[]Person{
		{FirstName:"One",LastName:"This",age:26},
		{FirstName:"One",LastName:"That",age:2},
		{FirstName:"Two",LastName:"Those",age:25},
	}
	fmt.Println(people)
	// the function passed to the slice is a closure
	sort.Slice(people,func(i,j int)bool{
		return people[i].LastName<people[j].LastName
	})
	fmt.Println(people)
	sort.Slice(people,func(i,j int)bool{
		return people[i].age<people[j].age
	})
	fmt.Println(people)
}

// passing functions as parameters to other functions is often useful
// when performing different operations on the same data.
