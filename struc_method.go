package main

import "fmt"

type Cust struct {
	Name, Addr string
	Age        int
}

func (cust Cust) sayHi() {
	fmt.Println("hello my name  is ", cust.Name)

}

func main() {
	cus := Cust{
		Name: "adetri",
		Addr: "jl teganga",
		Age:  18,
	}

	cus.sayHi()
}
