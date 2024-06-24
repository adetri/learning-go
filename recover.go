package main

import "fmt"

func defer_fun() {
	fmt.Println("defer execute")

	messaege := recover()
	fmt.Println("terjadi panic :", messaege)

}

func panic_fun(stat bool) {
	defer defer_fun()
	if stat {
		panic("panic execute bcs stat is true")
	}
}

func main() {

	panic_fun(true)
	fmt.Println("program is finis")
}
