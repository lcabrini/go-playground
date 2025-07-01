package main

import "fmt"

func main() {
	fmt.Print("    |")
	for col := 1; col <= 12; col += 1 {
		fmt.Printf("%4d|", col)
	}
	fmt.Println()

	for col := 0; col <= 12; col += 1 {
		fmt.Printf("----+")
	}
	fmt.Println()

	for row := 1; row <= 12; row += 1 {
		fmt.Printf("%4d|", row)

		for col := 1; col <= 12; col += 1 {
			fmt.Printf("%4d|", row*col)
		}
		fmt.Println()
	}
}
