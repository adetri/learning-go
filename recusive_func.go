package main

import "fmt"

func vactorial_loop(val int) int {
	result := 1

	for i := val; i > 0; i-- {
		result *= i
	}

	return result
}
func vactorial_loop1(val int) int {

	if val == 1 {
		return 1
	} else {
		return val * vactorial_loop1(val-1)
	}

}

func main() {

	fmt.Println(vactorial_loop1(10))

}
