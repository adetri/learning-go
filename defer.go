package main

import "fmt"

func test_defer() {

	fmt.Println("differ function")
}

func test() {

	defer test_defer()
	fmt.Println("run test fun")

}

func main() {

	test()

}
