package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)

	fmt.Println("随便输入点：")

	input, err := inputReader.ReadString('\n')

	if err == nil {
		fmt.Println("你输入的是:", input)
	} else {
		return
	}

	// 对unix:使用“\n”作为定界符，而window使用"\r\n"为定界符
	switch input {
	case "xiaoming\r\n":
		fmt.Println("welcome xiaoming")
	case "lilei\r\n":
		fmt.Println("welcome lilei!")
	default:
		fmt.Println("bye")
	}

	// version 2
	// 如果case带有fallthrough，程序会继续执行下一条case,不会再判断下一条case的expr
	switch input {
	case "xiaoming\r\n":
		fallthrough
	case "lilei\r\n":
		fmt.Printf("welcome %s\n", input)
	default:
		fmt.Println("bye")
	}

	//version 3

	switch input {
	case "xiaoming\r\n", "lilei\r\n":
		fmt.Printf("welcome %s\n", input)
	default:
		fmt.Println("bye!!")
	}
}
