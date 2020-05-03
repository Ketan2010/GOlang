//befor eritting our go code we define our 'main' package which is self defined
package main

//and we a import inbuilt package 'fmt' to use 'Println' method in it
import (
	"fmt"
)

//suppose we have two packages A and B
//methods/function in package  A: Aaa(), aaa(), a()
//methods in package B: Bbb(), B(), bbb(), b()
//now if we have to use methods of A in B then we can use only those methods which name starts with Capital letter
// methods Aaa() can be use in B
// methods Bbb(), B() can be use in A and not bbb(),b()
//methods which can be use in another package are called as exported name

//example:
//method Println() is in 'fmt' oackage but we can use it in our main package as it start with cpaital letter
//method toLarge() is also present in 'fmt' package but we cannot use it in our main
//hence any method that we have to use in another package must start with capital letter

func main() {
	fmt.Println("exported name")
}
