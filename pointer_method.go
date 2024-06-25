package main

import "fmt"

type Name struct {
	FirstName string
}

func (fn *Name) edit_name() {
	fn.FirstName = "Kampret"
}
func main() {
	det := Name{FirstName: "de"}

	det.edit_name()

	fmt.Println(det.FirstName)

}
