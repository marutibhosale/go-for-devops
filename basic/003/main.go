package main

import (
	"fmt"
)

func add(a,b int) (int) {
	add := a + b
	return add

}
func main() {
	addition := add(5,6)
	fmt.Println("addition is: ", addition)
}