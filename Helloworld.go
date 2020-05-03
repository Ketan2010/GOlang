//if package name is not main :go run: cannot run non-main package
package main

//fmt=format package should be use or import to perform printing
import (
	"fmt"
)

//if you dont write opening semicolon immiditely after main()
//then it will give error : semicolon missing or incomplete function body
func main() {
	fmt.Println("Hello world")

}
