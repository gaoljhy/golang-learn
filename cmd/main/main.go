/**
 * Copyright © 2020 Mr.G. All rights reserved.
 *
 * @author: Mr.G
 * @date: 2020-08-17
 * @TODO : {
	 1. 数组操作
	 2. 多维数组
	 3. 切片
	 4. 切片操作
	 :TODO}
*/
package main

import (
	"fmt"
)

func main() {
	ary := [...]int{1, 2, 3}
	var ary1 [3]int = [3]int{2, 4, 6}
	var ary3 [3][2]int = [3][2]int{{1, 2}, {1, 2}, {3, 4}}

	fmt.Println("length:", len(ary))
	fmt.Printf("type %T value%n \n", ary1, ary1)

	// 判断是否相等
	fmt.Println(ary == ary1)

	for _, v := range ary {
		fmt.Println(v)
	}
	// 输出多维数组中所有元素
	for i := 0; i < len(ary3); i++ {
		for j := 0; j < len(ary3[i]); j++ {
			fmt.Println(ary3[i][j])
		}
	}
	// 	切片使用

	var seq = ary1[:2]
	var newSeq []int
	var createSeq = make([]int, 5, 10)
	fmt.Println(newSeq == nil)
	fmt.Println(seq)
	fmt.Printf("make len : %d cap : %d\n", len(createSeq), cap(createSeq))

	// append 方法
	appSeq := append(newSeq, 2, 1)
	fmt.Println(appSeq)
	var copySeq []int = make([]int, 5)
	copy(copySeq, appSeq[:1])
	// copy 要看 第一个参数的 size
	fmt.Println(copySeq)
	// 前置位添加，也可后置位添加
	appSeq = append([]int{1, 2, 3}, copySeq[:2]...)
	// 不能用 数组来直接操作append，需要转换为 切片
	fmt.Println("0", appSeq)
	// 删除元素 第4个元素以后的元素
	appSeq = appSeq[:4]
	fmt.Println("1", appSeq)
	// 从中间第2,3个删除
	appSeq = append(appSeq[:1], appSeq[3:]...)
	for index, v := range appSeq {
		fmt.Println("index", index)
		fmt.Println("vaule", v)
	}
	fmt.Printf("len %d ,cap %d \n", len(appSeq), cap(appSeq))
	mutilVseq := [][]string{{"nihao", "world"}, {"hello", "world"}}
	fmt.Println(mutilVseq)
}
