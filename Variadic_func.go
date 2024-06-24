package main

import "fmt"

func sumall(num ...int) int {
	total := 0
	for _, number := range num {
		total += number
	}

	return total

}

func main() {

	fmt.Println(sumall(100, 200, 300, 400))

	number := []int{120, 80, 300, 400, 9000}
	fmt.Println(sumall(number...))

}
