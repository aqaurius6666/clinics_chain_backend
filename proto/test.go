package main

import "fmt"

type A interface {
	Test(n int)
}

type B struct {}

func (b B) Test(n int) {
	fmt.Println(n + 1)
}

type C struct {
	a A
}

func (c C) Test(n int) {
	fmt.Println(n + 1)
}

func main() {
	test := C{}
	test.a.Test(4)
}