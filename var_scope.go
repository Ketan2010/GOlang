package main

import (
	"fmt"
)

//there are two types of variabless
//1. function level:these are created inside function and can be access only inside that function
//2. package level : these are created at package level aand can be access by every functions in package

//package level
var pkglvl int = 50 //dont use undescore in variable names of go
var a int = 75

//you cant have two variables in same scope
func demo() {
	var funclvl int = 100
	//if we have same variable in function as that of package level variable
	var a int = 80
	fmt.Println("Function level:", funclvl)
	fmt.Println("Package level:", pkglvl)
	fmt.Println("Function level variable with same name of package level:", a) //will print 80
	//function will give more precedence to its own variables
}

func main() {
	demo()
	//fmt.Println("Function level:", funclvl) this willgive error
}
