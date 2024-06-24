package main

import "fmt"

type filtered func(string) bool

func validasiUsr(name string, filterz filtered) {

	if filterz(name) {
		fmt.Println("you're blocked")
	} else {
		fmt.Println("welcome ", name)
	}
}

func main() {
	filtered := func(name string) bool {
		return name == "anjing"
	}

	validasiUsr("anjing", filtered)

	validasiUsr("det", func(name string) bool {
		return name == "anjing"

	})
}
