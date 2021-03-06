package main

import (
	"fmt"
)

// 类型断言    了解空接口接受的值是什么类型
// 一种通过ok方式判断，另一种则用switch方法判断
func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	switch t := a.(type) {
	case string:
		fmt.Println("是一个字符串:", t)
	case int:
		fmt.Println("是一个int:", t)
	case int64:
		fmt.Println("是一个int64:", t)
	case bool:
		fmt.Println("是一个bool:", t)
	}
	/*str, ok := a.(string)
	if !ok {
		fmt.Println("猜错了")
	} else {
		fmt.Println("传进来的是一个字符串", str)
	}*/
}

func main() {
	assign(100)
	assign(true)
	assign(int64(200))
}
