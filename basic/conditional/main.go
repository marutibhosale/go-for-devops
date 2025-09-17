package main

import "fmt"

// select statement for concurrency with channels

var a = 3
var b = 4
func main() {
	if z := a + b; z > 5 {
		fmt.Println("z is greater than 5 and its value is", z)
	}else {
		fmt.Println("z is less than 5 and its value is", z)
	}
	
	
}