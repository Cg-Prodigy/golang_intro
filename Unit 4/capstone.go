package main

import (
	"fmt"
	"math/rand"
)
const(
	width=80
	height=15
)
type Universe [][]bool

func NewUniverse() Universe{
	u_verse:=make(Universe,height)
	for i:=range u_verse{
		u_verse[i]=make([]bool, width)
	}
	return u_verse
}

func (u Universe) Show(){
	for _,row:=range u{
		for _,col:=range row{
			if col==true{
				fmt.Printf("*")
			}else{
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
}
func (u Universe) Seed(){
	for i:=0;i<(width*height/4);i++{
		u[rand.Intn(height)][rand.Intn(width)]=true
	}
}

func (u Universe) Alive(x int,  y int ) bool{
	result:=false
	if y>len(u){
		y=len(u)%y
	}
	if x>len(u[y]){
		x=len(u[y])%x
	}
	if u[x][y]==true{
		result=true
	}
	return result
}
func (u Universe) Neighbors(x,y int) int{
	count:=0
	for i:=0;i<8;i++{
		if u.Alive(x,y+i)==true{
			count++
		}
	}
	return count
}
func (u Universe) Next(x,y int)bool{
	result:=false
	value:=u.Neighbors(x,y)
	if value<2{
		result=false
	} else if value==2 || value==3{
		result=true
	}else if value==3{
		result=true
	}else if value>3{
		result=false
	}
	return result

}
func main(){
	u_verse:=NewUniverse()
	u_verse.Seed()
	u_verse.Show()
	fmt.Println(u_verse.Alive(80,15))
}