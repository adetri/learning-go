package main

import "fmt"

func main() {

	name := "det"

	switch name {
	case "joko":
		fmt.Println("ini joko")
	default:
		fmt.Println(" this default value")

	}

	switch ln := len(name); ln < 4 {
	case true:
		fmt.Println("this switch case true condition")
	case false:
		fmt.Println("this switch case false condition")

	}
}
