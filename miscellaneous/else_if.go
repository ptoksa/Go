package main

import (
	"fmt"
)

func main() {
	a := 11
	b := 10
	if a < b {
		fmt.Println("a is less than b.")
	} else if a > b {
		fmt.Println("a is more than b.")
	} else {
		fmt.Println("a and b are equal.")
	}
}
