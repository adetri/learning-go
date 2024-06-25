package main

import (
	"errors"
	"fmt"
)

func pembagian(nilai int, pembagi int) (int, error) {

	if pembagi == 0 {
		return 0, errors.New("error pemabagi 0")
	} else {

		return nilai / pembagi, nil
	}
}

func main() {

	result, err := pembagian(100, 0)
	if err == nil {
		fmt.Println(result)
	} else {

		fmt.Println(err)
	}

}
