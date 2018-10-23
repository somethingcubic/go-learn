package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	a := Person{
		Name: "joe",
		Age:  19,
	}
	fmt.Println(a)
}
