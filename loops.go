package main

import "fmt"

func main() {
	//in go language only for loop is exist
	//this will print nifinite number of times
	//for {
	//	fmt.Println("Hello go!")
	//}
	//first way of writing for loop
	var i int = 0 //i:= 0  also works
	for i <= 4 {
		fmt.Println("Hello go!")
		i++
	}

	//second way of writing for loop
	for i := 1; i <= 5; i++ {
		fmt.Println("for is king of GO ")
	}

}
