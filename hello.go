package main

import "fmt"

var x int
var y string
var z bool

func main() {
	x = 42
	y = "James Bond"
	z = true
	s := fmt.Sprintf("x is %d, y is %s, and z is %t", x, y, z)
	fmt.Println(s)
}
