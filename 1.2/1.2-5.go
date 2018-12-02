package main

import "fmt"

func main() {
	var pI *int // memory address ==> of a value of type int

	var I int = 3
	pI = &I // pI ==> points to the location of variable I
	increment(pI)
	fmt.Println(*pI)
}

func increment(pI *int) {
	*pI++ // dereferencing
}
