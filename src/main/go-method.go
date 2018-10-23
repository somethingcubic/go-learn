package main

import "fmt"

type A int

func main() {
	var a A
	a.Increase(100)
	fmt.Println(a)
}

func (a *A) Increase(num int) {
	*a += A(num)
}
