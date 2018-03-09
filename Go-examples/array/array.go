/*

	在Go语言中，数组是值类型和长度是类型的组合
	也就是说"[10]int"和“[20]int”是完全不同的两种数组类型
	同类型的两个数组支持"=="和"!="比较，但是不能比较大小
	数组作为参数时，函数内部不改变数组内部的值，除非是传入数组的指针。
	数组的指针：*[3]int
	指针数组：[2]*int

*/

package main

import "fmt"

func main() {
	//创建有5个元素的整型数组
	var a [5]int
	fmt.Println("empty:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	//定义并初始化数组
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

}
