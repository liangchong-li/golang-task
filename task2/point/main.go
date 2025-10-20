package main

import (
	"fmt"
)

func main() {
	num1 := 2
	fmt.Println("原始变量num1:", num1)
	point1(&num1)
	fmt.Println("修改后变量num1:", num1)

	var slide []int = []int{1, 2, 3}
	fmt.Println("原始变量slide:", slide)
	point2(&slide)
	fmt.Println("修改后变量slide:", slide)
}

// 题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，
// 然后在主函数中调用该函数并输出修改后的值。
// 考察点 ：指针的使用、值传递与引用传递的区别。
func point1(numP *int) {
	numPP := &numP
	**numPP = *numP + 10
}

// 题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
// 考察点 ：指针运算、切片操作。
func point2(slidep *[]int) {
	slide := *slidep
	for i, _ := range slide {
		slide[i] *= 2
	}
}
