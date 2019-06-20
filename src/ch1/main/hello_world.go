package main

import (
	"fmt"
	"os"
)

func main() {
	//fmt.Println("Hello World")

	if len(os.Args) > 1 {
		fmt.Println("Hello World => ", os.Args[1])
	}

	os.Exit(0)
}
