package main

import "fmt"

 var a = 2

func main() {
	b := 4
	x := 8
	y := 16
	fmt.Printf("%d \t %b\n", 1,1)
	fmt.Printf("%d \t %b\n", 1<<1, 1<<1)
	fmt.Printf("%d \t %b\n", 2<<2, 1<<2)

	fmt.Printf("%v type of %T\n", a ,a)
	fmt.Printf("%v type of %T\n", b ,b)

	fmt.Printf("value of x is %v and its type is %T\n", x, x)
	fmt.Printf("value of y is %v and its type is %T\n", y, y)

}