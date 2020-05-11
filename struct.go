//structure is user defind data type in go
package main

import (
	"fmt"
)

type Student struct {
	rollno int
	name   string
	marks  int
}

func main() {
	var s1 = Student{22, "ketan", 50}
	fmt.Println(s1)
	var s2 = Student{rollno: 2, marks: 90}
	fmt.Println(s2)

}
