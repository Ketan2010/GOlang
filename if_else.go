package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Println("enter the number:")
	fmt.Scanf("%d", &num)
	if num%2 == 0 {
		fmt.Printf("%d is even number.\n", num)
	} else {
		fmt.Printf("%d is odd number.\n", num)
	}

}
