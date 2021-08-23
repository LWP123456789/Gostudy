package main

import "fmt"

// 匿名字段
// 使用于字段比较少比较简单的场景，不常用

type person struct {
	string
	int
}

func main() {
	p1 := person{
		"lep",
		20,
	}
	fmt.Println(p1.string)
}
