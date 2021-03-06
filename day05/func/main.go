package main

import "fmt"

// 函数

// 函数存在的意义？
// 函数是一段代码的封装
// 把一段逻辑抽象出来封装到一个函数中，给它起个名字，每次用到它的时候直接使用函数名调用即可
// 使用函数能够让代码结构更清晰、更简洁

// 函数的定义
func sum(x int, y int) (ret int) {
	return x + y
}

// 没有返回值
func f1(x int, y int) {
	fmt.Println(x + y)
}

// 没有参数
func f2() {
	fmt.Println("f2")
}

// 没有参数但由返回值的
func f3() int {
	return 3
}

// 多个返回值
func f4() (int, string) {
	return 1, "lep101"
}

// 参数的类型简写：
// 当参数中连续多个参数的类型一致时，我们可以将非最后一个参数的类型省略
func f5(x, y, z int, m, n, string, i, j bool) int {
	return x + y
}

// 可变长参数
// 可变长参数必须放在函数参数的最后
func f6(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y) // y的类型是切片 []int
}

// Go语言中函数没有默认参数这个概念 提倡所见即所得

// 参数可以命名也可以不命名 起名了可以在函数内直接使用
func main() {
	fmt.Println(sum(10, 20))
	_, n := f4()
	fmt.Println(n)
	f6("下雨了")
	f6("下雨了", 1)
	f6("下雨了", 1, 2, 3, 4, 5, 6, 7)
}
