package main

import "fmt"

func Hallow() {
	fmt.Println("hallow")
}

func PrintString(data string) {

	fmt.Println(data)
}

func Retfunc(data string) string {
	return data
}

func MultipleReturn() (string, string) {

	return "ade", "tri"
}

func main() {

	Hallow()
	PrintString("param is param")

	retfun := Retfunc("this ret")
	fmt.Println(retfun)

	firstName, lastName := MultipleReturn()

	fmt.Println(firstName, lastName)

	firstName1, _ := MultipleReturn()

	fmt.Println(firstName1)

}
