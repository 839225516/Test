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
