package main

import "fmt"

type Cust struct {
	Name, Addr string
	Age        int
}

func main() {

	var det Cust

	det.Name = "dett"
	det.Addr = "jl kecai"

	fmt.Println(det)

	kampret := Cust{
		Name: "kampret",
		Addr: "jl.cabul",
		Age:  19,
	}

	fmt.Println(kampret)

	kampak := Cust{"kampank", "jl.cabs", 19}

	fmt.Println(kampak)
}
