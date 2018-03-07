package main

import "fmt"
import "math"

//Go支持定义字符常量，字符串常量，布尔型常量和数值常量
const s string = "constant"

func main() {
	fmt.Println(s)

	// 数值型常量没有具体类型，除非指定一个类型
	const n = 50000

	const d = 3e20 / n
	fmt.Println(d)

	// 显式类型转换
	fmt.Println(int64(d))

	//n转换成float64
	fmt.Println(math.Sin(n))
}
