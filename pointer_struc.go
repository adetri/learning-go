package main

import "fmt"

type Address struct {
	City string
}

func editCity(addr *Address) {
	addr.City = "Jakarta Selatan"
}

func main() {

	addr := Address{
		City: "jakarta",
	}

	editCity(&addr)

	fmt.Println(addr)

}
