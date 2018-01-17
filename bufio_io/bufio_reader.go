// bufio.NewReader（）
// func NewReader(rd io.Reader) *Reader

// ReadString
// func (b *Reader) ReadString(delim byte) (line string, err error)

// inputReader 是个指针，它指向一个 bufio类的Reader,
// main函数里，通过了bufio.NewReader(os.Stdin)创建了一个buffer reader,
// 并联接到inputReader这个变量

package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	inputReader *bufio.Reader
	input       string
	err         error
)

func main() {
	inputReader = bufio.NewReader(os.Stdin)

	fmt.Println("随便输入点：(以S结尾)")

	// 'S' 这里使用S表示结束符，也可以用其它，如'\n'
	input, err := inputReader.ReadString('S')

	if err != nil {
		fmt.Println("有一个错误", err)
	} else {
		//去除结尾的S
		input = input[:len(input)-1]
		fmt.Printf("你输入的是：%s\n", input)
	}

}
