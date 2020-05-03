package main

import (
	"fmt"
)

func main() {
	//in case of c,java,python if we declare variable and we dont use it in future then it will not give erro
	//but GO will give error if we dont use declared variable or packages:declared but not used
	//ariable declaration: var <name> <type>=value
	var num1 int = 9
	var num2 int = 3
	fmt.Println("addition:", num1+num2)
	//only declaraton: var num int : its default value will be 0
	var x, y int
	x, y = 3, 1
	var z int = x * y
	fmt.Println("Multiplication:", z)
	result := 22 // this is similler to :var result =22
	fmt.Println("result:", result)
	const k = 9.78
	//k = 90: error: cannot assign to k (as it is a constant)
	fmt.Print("constant:", k)
}
