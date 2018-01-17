// Scanln 将从标准输入的带有空格的字符串值保存到相应的变量里去，并以一个换行结束输入，
// Scanf做相同的工作，但它使用第一个参数指时输入格式，
// Sscan系列函数也是读取输入，但它是用来从字符串变量里读取，而不是从标准（os.Stdin）里读取

package main

import (
	"fmt"
)

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "12.12 / 4232 / go"
	format                 = "%f / %d / %s"
)

func main() {
	fmt.Println("请输入名字")
	fmt.Scanln(&firstName, &lastName)
	//fmt.Scanf("%s %s",&firstName,&lastName)
	fmt.Printf("Hi,%s %s!\n", firstName, lastName)
	fmt.Sscanf(input, format, &f, &i, &s)
	fmt.Println("From the string we read:", f, i, s)

}
