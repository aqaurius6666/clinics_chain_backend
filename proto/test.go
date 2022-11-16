package main

import "fmt"

type B struct {
	g string
}

type A struct {
	a string
	b int
	s B
}

func main() {
	a := 5
	b, ok := interface{}(a).(int)
	fmt.Println(b, ok)
}
