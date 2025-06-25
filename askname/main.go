package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	fmt.Print("What is your name? ")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read from stdin: %s", err)
		os.Exit(1)
	}

	name = strings.TrimSpace(name)
	caser := cases.Title(language.English)
	if name == "" {
		fmt.Println("Coward!")
	} else {
		fmt.Printf("Hi, %s and welcome to Go!\n", caser.String(name))
	}
}
