package main

import "fmt"

func main() {
	var x int
	fmt.Printf("Print numbers:")
	fmt.Scan(&x)
	y := 2
	for x >= y {
		switch {
		case x%y == 0:
			fmt.Println(y)
			x = x / y
		case x%y != 0:
			y += 1
		default:
			break
		}
	}

}
