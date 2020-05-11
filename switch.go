package main

import (
	"fmt"
)

func main() {

	var ch int
	var n1 int = 6
	var n2 int = 3

	fmt.Printf("\n============MENU==============")
	fmt.Printf("\n1.ADD\n 2.SUB \n 3.MUL \n 4.DIV \n")
	fmt.Println("enter choice:")
	fmt.Scanf("%d", &ch)
	switch ch {
	case 1:
		fmt.Printf("\nADDITION:%d", n1+n2)
		break
	case 2:
		fmt.Printf("\nSUBTRACTION:%d", n1-n2)
		break
	case 3:
		fmt.Printf("\nMULTIPLICATION:%d", n1*n2)
		break
	case 4:
		fmt.Printf("\nDIVISION:%d", n1*n2)
		break
	default:
		fmt.Printf("\nWRONG INPUT!!!!")

	}

}
