package main

import "fmt"

func main() {
	counter := 0
	increment := func() {
		fmt.Println("incremnet execute")

		counter++
	}

	increment()
	increment()

	fmt.Println(counter)

}
