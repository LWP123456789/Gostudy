package main

import "fmt"

// panic 和 recover
// 例如程序一开始连接数据库 没链接成功此时应用panic

func funcA() {
	fmt.Println("a")
}

func funcB() {
	// 刚刚打开了数据库连接
	defer func() {
		err := recover() //recover会恢复程序错误之前的状态，使程序继续执行
		fmt.Println(err)
		fmt.Println("释放数据库连接...")
	}() // 在panic之前释放数据库连接
	panic("出错啦！！！")  // 抛出了一个严重的错误 程序崩溃退出
	fmt.Println("b") // Unreachable code panic后永远到不了的代码(起不了作用)
}

func funcC() {
	fmt.Println("c")
}

func main() {
	funcA()
	funcB()
	funcC()
}
