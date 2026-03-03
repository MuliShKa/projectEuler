package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Printf("Enter number: ")
	fmt.Scan(&num)
	sum := 0
	for i := 0; i < num; i++ {
		if i%5 == 0 || i%3 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
