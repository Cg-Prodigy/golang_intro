package main

import (
	"fmt"
	"strings"
)


var t interface{
	talk() string
}

type martian struct{}
type laser int
type planet struct{}

func (m martian) talk()string{
	return "nack nack"
}
func(l laser) talk()string{
	return strings.Repeat("pew",int(l))
}

func main(){
	t=martian{}
	fmt.Println(t.talk())
	t=laser(3)
	fmt.Println(t.talk())

}