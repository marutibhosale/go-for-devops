package main

import (
	"flag"
	"fmt"
	"reflect"
)

func main() {
	// Define flags
	verbose := flag.Bool("v", false, "Enable verbose mode")
	name := flag.String("name", "World", "A name to say hello to")
	age := flag.Int("age", 0, "Your age")
    fmt.Println(reflect.TypeOf(verbose))
	test := &verbose
	fmt.Println(**test)
	fmt.Println(&verbose)
	// Parse the flags
	flag.Parse()

	// Use the flags
	if *verbose {
		fmt.Printf("Verbose mode enabled\n")
	}
	fmt.Printf("Hello, %s!\n", *name)
	fmt.Printf("Your age is %d\n", *age)
}