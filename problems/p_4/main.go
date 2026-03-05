package main

import (
	"fmt"
)

func main() {
	polidrom := 0
	temp := 0
	var best int
	for a := 999; a > 100; a-- {
		for b := 999; b > 100; b-- {
			polidrom = a * b
			temp = polidrom

			n1 := temp % 10
			temp /= 10
			n2 := temp % 10
			temp /= 10
			n3 := temp % 10
			temp /= 10
			n4 := temp % 10
			temp /= 10
			n5 := temp % 10
			temp /= 10
			n6 := temp % 10

			if n1 == n6 && n2 == n5 && n3 == n4 {
				if polidrom > best {

					best = polidrom
				}
			}

		}

	}
	fmt.Println(best)
}
