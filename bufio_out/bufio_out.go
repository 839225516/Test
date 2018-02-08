package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	w := bufio.NewWriter(os.Stdout)

	fmt.Fprint(w, "Hello,")
	fmt.Fprint(w, "world!\n")

	//别忘记释放
	w.Flush()
}
