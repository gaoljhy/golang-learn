package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("hello world")
	//  input flag class
	input := flag.String("input", "default", "input some string in this param")
	flag.Parse()
	// print
	fmt.Println("input content:", *input)

}
