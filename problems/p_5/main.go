package main

import "fmt"

func main() {
	n := 20

	for {
		ok := true

		for i := 1; i <= 20; i++ {
			if n%i != 0 {
				ok = false
				break
			}
		}

		if ok {
			fmt.Println(n)
			return
		}

		n += 20
	}
}
