package main

import (
	"fmt"
)

type Filt func(string) string

func paramfunc(name string, filter Filt) {
	filtername := filter(name)
	fmt.Println(filtername)

}

func filtered(name string) string {

	if name == "anjing" {
		return "******"
	} else {

		return name
	}
}

func main() {
	paramfunc("anjing", filtered)

}
