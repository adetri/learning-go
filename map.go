package main

import (
	"fmt"
)

func main() {
	name := map[string]string{

		"name":  "test1",
		"addrs": "addrs2",
	}

	fmt.Println(name)
	delete(name, "name")
	fmt.Println(name)

	book := make(map[string]string)
	book["title"] = "test"
	book["author"] = "author name"

	fmt.Println(book)
}
