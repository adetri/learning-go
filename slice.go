package main

import "fmt"

func main() {
	name_array := [...]string{"ade", "tri", "jumli", "vanhouten"}
	name_slice := name_array[1:3]
	name_slice = append(name_slice, "tesst", "asdasd", "asdsas")
	fmt.Println(name_array)
	fmt.Println(name_slice)
	fmt.Println(len(name_array))
	fmt.Println(len(name_slice))
	fmt.Println(cap(name_array))
}
