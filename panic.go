package main

import "fmt"

func defer_func() {
	fmt.Println("Defer Called")
}

func test_panic_fun(status bool) {
	defer defer_func()
	if status {
		panic("error")
	}

}

func main() {
	test_panic_fun(true)
}
