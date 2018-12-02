package main

import "fmt"

func main() {
	defer fmt.Println("World 1")
	defer fmt.Println("world 2")
	fmt.Println("Hello")
	testpanics()
	fmt.Println("world")
}

func testpanics() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			fmt.Println("We recovered from a panic!!")
		}
	}()
	panic("A panic happende")
}
