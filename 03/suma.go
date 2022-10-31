package main

import (
	"fmt"
	"Users/Valentin/Documents/Go/examplegocode/01-test/03/acdc"
)
func main() {

	fmt.Println("2 + 3 = ", acdc.Sum(2,3))
	fmt.Println("3 + 7 = ", acdc.Sum(2,3,4,5,6,7,8,9))
}