//when we write keyword defer before any function
//then that function will be only evaluated and after execution of other functions it will be executed
//example
// func a(){
// 	fmt.println("in a")
// 	defer b()
// 	fmt.println("end a")
// }
// func b(){
// 	fmt.println("in b")
// }
//above code output will be
//in a
//end a
//in b
package main

import (
	"fmt"
)

//if there are multiple defer functions then after only evaluation these are kept in stack
//and after exacution of all other functions defer functions in stack will kept out one by one for rxecution
func main() {
	for i := 1; i <= 6; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("bye")

}

// OUTPUT
// bye
// 6
// 5
// 4
// 3
// 2
// 1
