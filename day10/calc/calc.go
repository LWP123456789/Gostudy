package calc

import "fmt"

// Add 包中的标识符(变量名\函数名\结构体\接口等)如果首字母是小写的,表示私有(只能在当前这个包中使用)
// 首字母大写的标识符可以对外部的包可见

func init() {
	fmt.Println("自动执行!")
}

func Add(x, y int) int {
	return x + y
}
