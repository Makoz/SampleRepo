package main

import (
	"fmt"
	"os"
	// "os/exec"
)

func func1() {
	fmt.Println("func1 from test2")
}

func func2() {
	fmt.Println("func2 from test2")
}

func main() {

	usage := fmt.Sprintf("Usage: %s CF Id\n", os.Args[0])
	if len(os.Args) != 2 {
		fmt.Printf(usage)
		os.Exit(1)
	}

	cfId := os.Args[1]
	if cfId == "1" {
		func1()
	} else {
		func2()
	}

}
