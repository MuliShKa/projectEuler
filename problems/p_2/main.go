package main

import "fmt"

func main() {
	var endNum int
	fmt.Print("Enter end num: ")
	fmt.Scan(&endNum)

	sum := 0
	n1 := 0
	n2 := 1

	for n1 < endNum {
		if n1%2 == 0 {
			sum += n1
		}
		fmt.Println(n2)
		fibonacci := n1 + n2
		n1 = n2
		n2 = fibonacci
	}

	fmt.Println("Sum:", sum)
}
