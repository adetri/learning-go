package main

import "fmt"

func main() {

	type NoKtp string

	var nik string = "test1"
	var true_nik NoKtp = NoKtp(nik)

	fmt.Println(true_nik)
}
