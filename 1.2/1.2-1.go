package main

import "fmt"

func main() {
	fmt.Println(CompareNumbers(1, 2))

	switch x := 2; x {
	case 3:
		fmt.Println("I am 3")
	case 2:
		fmt.Println("I am 2")
		// fallthrough
	case 4:
		fmt.Println("I am 4")
	case 5:
		fmt.Println("I am 5")
	}

	i := 0
	for i <= 10 {
		fmt.Println(i)
		i++
	}
}

func CompareNumbers(i1, i2 int) (bool, int) {
	/*
		if i1 > i2 {
			fmt.Println("First number is greater than the second number")
			return false, i1 - i2
		} else if i2 > i1 {
			fmt.Println("Second number is greater than the first number")
			return false, i2 - i1
		}
	*/

	switch {
	case i1 > i2:
		fmt.Println("First number is greater than the second number")
		return false, i1 - i2
	case i2 > i1:
		fmt.Println("second numbrer is greater than the first number")
		return false, i2 - i1
	}

	fmt.Println("numbers are equal!!")
	return true, 0
}
