package main

import (
	"fmt"
	"os"
)



func main(){


}

func init() {
	var user = os.Getenv("USER")
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("recovery",x)
		}
	}()
	if user != "" {
		panic("have value for $USER")
	} else {
		fmt.Println(user)
	}

}
