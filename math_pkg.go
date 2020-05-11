//
package main

import (
	"fmt"
	"math"
)

func main() {

	var k float64 = 12.0
	var result float64 = math.Sqrt(k)
	//fmt.Println("square root:", result) this will give  3.4641016151377544
	//to frint only upto .2 decimalss
	//make sure that at the time of formating (%f,%g,etc) use Printf function
	fmt.Printf("square root of %g is %.2f\n", k, result)
	//round function
	fmt.Printf("Rounding:%f \n", math.Round(3.14))
	//ceil function
	fmt.Printf("Ceiling:%f\n", math.Ceil(3.56))
	//floor function
	fmt.Printf("flooring:%f", math.Floor(3.56))

}
