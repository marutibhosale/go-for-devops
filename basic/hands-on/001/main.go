package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i <5; i++ {
		x := rand.Intn(5)
		fmt.Println(x)

		// switch x {
		// case 0:
		// 	fmt.Println("valune of x is 0")
		// case 1:
		// 	fmt.Println("valune of x 1")
		// case 2:
		// 	fmt.Println("valune of x is 0")
		// case 3:
		// 	fmt.Println("valune of x 1")
		// case 4:
		// 	fmt.Println("valune of x is 0")
		// }
	}

}