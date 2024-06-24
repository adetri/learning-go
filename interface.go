package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func sayHi(value HasName) {
	fmt.Println(value.GetName())
}

func main() {
	per :=
		Person{Name: "det"}

	sayHi(per)

}
