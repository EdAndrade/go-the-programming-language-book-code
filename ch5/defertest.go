package main

import (
	"fmt"
	"os"
)

func main() {
	deferTesting := deferTest()
	resp, err := deferTesting.Open()
	if err != nil {
		fmt.Errorf("Someting happen %v", err)
		os.Exit(1)
	}
	defer deferTesting.Close()
	fmt.Println(resp)
	deferTesting.Execute()
}

type deferReturn struct {
	Close   func()
	Open    func() (status string, err error)
	Execute func()
}

func deferTest() deferReturn {
	Close := func() {
		fmt.Println("Closing")
	}

	Open := func() (status string, err error) {
		return "Opened", nil
	}

	Execute := func() {
		fmt.Println("Executing")
	}

	return deferReturn{
		Close, Open, Execute,
	}
}
