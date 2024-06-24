package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {

		if i == 10 {
			fmt.Println("break at :", i)
			break
		}

	}

	fmt.Println("this end")
}
