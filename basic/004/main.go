package main

import "fmt"

func swap(x , y string) (string, string) {
	return y, x
}
const test int = 10 
func main(){
	a, b := swap("Hi", "Hello")
	fmt.Println(a, b)
}