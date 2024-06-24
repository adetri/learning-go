package main

import "fmt"

func main() {

	counter := 0

	for i := 0; i < 10; i++ {
		counter = i
		fmt.Println(counter)
	}

	name := []string{"ade", "tri", "jumli"}

	for i := 0; i < len(name); i++ {

		fmt.Println(name[i])

	}

	for idx, name1 := range name {
		fmt.Println("idx", idx, "name", name1)
	}
}
