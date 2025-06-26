package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Name of the executable: %s\n", os.Args[0])

	for i, a := range os.Args[1:] {
		fmt.Printf("%2d. %s\n", i, a)
	}
}
