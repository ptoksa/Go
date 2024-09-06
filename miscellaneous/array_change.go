package main

import (
	"fmt"
)

func main() {
	prices := [3]int{10, 20, 30}

	prices[0] = 50
	fmt.Println(prices)
}
