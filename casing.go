package main

import (
	"fmt"
	"reflect"
)

func random() any {

	return "OK"
}

func main() {

	result := random()
	fmt.Println(result)
	fmt.Println(reflect.TypeOf(result))

}
