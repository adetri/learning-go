package main

import "fmt"

func main() {

	var name [3]string
	name[0] = "sia pa"
	name[1] = "dia"
	name[2] = "disana"

	fmt.Println(name)

	var name2 = [3]string{
		"dimana", "dia", "berada",
	}

	fmt.Println(name2)

	var name3 = [...]string{}

	fmt.Println(name3)
	fmt.Println(len(name3))

}
