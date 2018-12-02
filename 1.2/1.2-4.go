package main

import "fmt"

type person struct {
	Name    string
	Age     int
	Address string
}

func main() {
	json := person{
		Name:    "Json S.",
		Age:     38,
		Address: "Germany",
	}
	fmt.Println(json.Name)
}
