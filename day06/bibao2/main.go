package main

import "fmt"

// 闭包是什么？
// 闭包是一个函数，这个函数包含了他外部作用域的一个变量

// 底层的原理：
// 1、函数可以作为返回值
// 2、函数内部查找变量的顺序，现在自己内部找，找不到往外层找

// 闭包返回值是一个函数 在返回函数的外头的变量是不会在函数运行完后立即被销毁的
// 相当于开辟了一个内存存放参数，而这个参数值会被返回函数所影响

func adder(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func main() {
	ret := adder(320)
	fmt.Println(ret(200))
	ret2 := ret(200)
	fmt.Println(ret2)
}
