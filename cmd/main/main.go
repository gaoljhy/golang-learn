/**
 * Copyright © 2020 Mr.G. All rights reserved.
 *
 * @author: Mr.G
 * @date: 2020-08-17
 * @TODO : {
	 1. 类型转换
	 :TODO}
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "123455"
	strtoi, _ := strconv.Atoi(str)
	strtoi++
	fmt.Printf("type %T value%n \n", str, strtoi)

}
