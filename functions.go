package main

import (
	"fmt"
)

//func <name>(para1 <type>,para2 <type>) <type_of_return> {<body>}
func add(x int, y int) int {
	var z = x + y
	return z
}
func calc(x int, y int) (int, int) {
	var m = x * y
	var s = x - y
	return m, s

}
func main() {
	var n1 int = 4
	var n2 int = 5
	var result = add(n1, n2)
	var mul, sub int
	mul, sub = calc(n1, n2)
	fmt.Println("Addition:", result)
	fmt.Println("Multiplication:", mul)
	fmt.Println("Subtraction", sub)
}
