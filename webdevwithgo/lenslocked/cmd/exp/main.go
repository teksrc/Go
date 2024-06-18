package main

import (
	"fmt"
)

func blah() {
	panic("...")
}

func main() {
	fmt.Println("Experimental: ")
	blah()
}
