package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println("Value of i is", i)
	}

	// y := 1
	// for { fmt.Println("print", y); y++ if y > 10 { break } }
	y := 25
	for y > 20 { // while loop
		fmt.Println("Value of y is", y)
		y--
	}

	for i := range 10 {
		fmt.Println("Value of i is", i)
	}

	l := []int{10, 20, 30, 40, 50}
	for i, v := range l {
		fmt.Printf("index is %v and value is %v\n", i ,v)
	}

	m := map[string]int { "test":1, "test1":2}

	for k, v := range m {
		fmt.Printf("key %v has value %v\n", k, v )
	}

	for pos,ch := range "maruti" {
		fmt.Printf("postion %v has charater %v\n", pos, ch)
	}

	for i := 0; i <=100; i++ {
		fmt.Println(i)
	}
}
