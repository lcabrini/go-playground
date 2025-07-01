package main

import "fmt"

func main() {
	a := 5
	b := 7

	fmt.Printf("Before swap, a = %d, b = %d\n", a, b)
	swap(&a, &b)
	fmt.Printf("After swap,  a = %d, b = %d\n", a, b)
}

func swap(i *int, j *int) {
	*i, *j = *j, *i
}
