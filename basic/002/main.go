package main

import (
	"fmt"
	"math/rand"
)

func main () {
	fmt.Println("random no: ", rand.Intn(100))
	fmt.Println("random no: ", rand.Intn(100))
	fmt.Println("random no: ", rand.Intn(100))
	fmt.Println("random no: ", rand.Intn(100))
	fmt.Println("random no: ", rand.Intn(100))
	fmt.Println("random no: ", rand.Intn(100))
	fmt.Println("random no: ", rand.Intn(100))

	fmt.Printf("random  test %v and %T\n", rand.Intn(100),rand.Intn(100))
}