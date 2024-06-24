package main

import "fmt"

func main() {

	name := "det"

	if name == "det" {

		fmt.Print("ini ")
		fmt.Println(name)

	} else if name == "joko" {
		fmt.Println("jokoman nihh")
	} else {
		fmt.Println(" undifind")
	}

	if ln := len(name); ln < 4 {
		fmt.Println(ln)
		fmt.Println("nama nya dia kurang dari 4")
	}
}
