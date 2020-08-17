package main

import (
	"fmt"
)

//  重新定义一个 type
type num int

const (
	one num = iota
	two
	three
	four
	five
)
const ten int32 = 10

func main() {

	fmt.Println(one)
	fmt.Println(ten)

}
