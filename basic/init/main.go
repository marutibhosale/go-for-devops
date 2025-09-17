package main

import "fmt"

func init() { // first run init by default without calling
	fmt.Println("Printing from init function")
}

func main() {
	fmt.Println("Printing from main function")
}

func test() {
	fmt.Println("Printing from test function")
}
