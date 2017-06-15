package main

import (
	"fmt"
)

func main() {
	s := make([]int, 10)
	fmt.Printf("Addr of first element: %p, Type: %T\n", s[0], s[0])
	fmt.Printf("Addr of first element: %p, Type: %T\n", &s[0], &s[0])

	fmt.Printf("Addr of slice itself:  %p, Type: %T\n", s, s)
	fmt.Printf("Addr of slice itself:  %p, Type: %T\n", &s, &s)

	// Addr of first element: %!p(int=0), Type: int
	// Addr of first element: 0x10432180, Type: *int
	// Addr of slice itself:  0x10432180, Type: []int
	// Addr of slice itself:  0x1040a130, Type: *[]int

}
