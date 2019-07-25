package main

import "fmt"

type NUMBER int

var x NUMBER

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
