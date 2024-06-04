package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[:], " ")
	for index, arg := range os.Args[1:] {
		fmt.Println(index, arg)
	}
}