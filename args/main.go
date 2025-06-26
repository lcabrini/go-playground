package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	name := filepath.Base(os.Args[0])
	fmt.Printf("Name of the executable: %s\n", name)

	for i, a := range os.Args[1:] {
		fmt.Printf("%2d. %s\n", i, a)
	}
}
