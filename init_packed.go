package main

import (
	"fmt"
	"test_project/db"
	_ "test_project/test_init"
)

func main() {

	fmt.Println(db.GetDbName())

}
