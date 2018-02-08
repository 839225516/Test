package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
)

func Md5SumFile(file string) (value [md5.Size]byte, err error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	value = md5.Sum(data)
	return
}

func main() {
	v_name := "account_confirm"
	v_md5 := "c16a800b4a97600e5a0157b134f7836d"
	v_version := "v1.0"

	md5Value, err := Md5SumFile(v_name)
	if err != nil {
		fmt.Println(v_name, "文件不存在\n")
	} else {
		//fmt.Println(md5Value)
		//将[]byte转成16进制
		md5str := fmt.Sprintf("%x", md5Value)
		//fmt.Println(md5str)
		if v_md5 == md5str {
			fmt.Println(v_name, "的版本是:", v_version)
		} else {
			fmt.Println("版本不一致")
		}
	}
}
