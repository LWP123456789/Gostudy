package main

import "fmt"

// 自定义类型和类型别名

// type后面跟的是类型
type myInt int
type yourInt = int // 类型别名

func main() {
	var n myInt = 100
	var m yourInt = 10
	fmt.Println(n)
	fmt.Printf("%T\n", n)
	fmt.Println(m)
	fmt.Printf("%T\n", m)

	var c rune
	c = '中'
	fmt.Println(c)
	fmt.Printf("%T\n", c)
}
