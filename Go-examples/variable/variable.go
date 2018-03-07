/*
************************************
Go 的基本类型有：

bool
string
int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr
byte // uint8 的别名
rune // int32 的别名 代表一个Unicode码
float32 float64
complex64 complex128

*************************************
*/

package main

import "fmt"

func main() {
	var a string = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	//定义变量时，没有给出初始值的整数被默认初始化为零值
	var e int
	fmt.Println(e)

	// ":=" 是同时定义和初始化变量的快捷方式
	f := "short"
	fmt.Println(f)
}
