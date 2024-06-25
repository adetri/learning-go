package main

import "fmt"

type validationErr struct {
	Msg string
}

func (v validationErr) Error() string {

	return v.Msg

}

type notFoundErr struct {
	Msg string
}

func (n notFoundErr) Error() string {

	return n.Msg

}

func saveData(id string, data any) error {

	if id == "" {
		return &validationErr{Msg: "This validation Error"}
	}

	if id != "det" {
		return &notFoundErr{Msg: "Id not found"}
	}

	return nil

}

func main() {

	err := saveData("det", "")
	if err != nil {
		if validationErr, ok := err.(*validationErr); ok {
			fmt.Println("validation err", validationErr.Msg)
		} else if notFoundErr, ok := err.(*notFoundErr); ok {
			fmt.Println("id not found", notFoundErr.Msg)
		} else {
			fmt.Println("Unknow err")
		}

		switch finalerr := err.(type) {
		case *validationErr:
			fmt.Println("validation err", finalerr.Error())
		case *notFoundErr:
			fmt.Println("not found id", finalerr.Error())
		default:
			fmt.Println("unknow err")
		}
	} else {
		fmt.Println("success to insert")
	}

	fmt.Println("this execute to")

}
