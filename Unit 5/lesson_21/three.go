package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type create_user struct {
	F_name, L_name, Email string
}	
func main() {
	a_user := create_user{"Ego","Brian", "ego@go.com"}
	bytes,err:= json.Marshal(a_user)
	errorCheck(err)
	fmt.Println(string(bytes))
	
}

// utilities
func errorCheck(err error){
	if err!=nil{
		fmt.Println(err)
		os.Exit(1)
	}
}
