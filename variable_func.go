package main

import "fmt"

func Goodbye(data string) string {

	return "Goodbay " + data
}

func main() {
	gb := Goodbye

	fmt.Println(gb("mael"))

}
