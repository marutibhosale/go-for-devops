package main

import (
	"fmt"
)

func add(a int,b int) (int) {
	add := a + b
	return add

}
func main() {
	fmt.Println(add(4,5))
	hello()
}

func hello() {
	fmt.Println("Hello")
}