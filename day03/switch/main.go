package main

import "fmt"

//switch
//简化大量的判断
//不用break

func main()  {
	var n = 5

	switch n {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的数字")
	}

	switch n := 7; n {
	case 1,3,5,7,9:
		fmt.Println("奇数")
	case 2,4,6,8:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}

	//分支可以使用表达式
	var age = 30
	switch {
	case age < 25:
		fmt.Println("好好学习吧")
	case age > 25 && age < 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")
	}

	//fallthrough语法可以执行满足条件的case的下一个case，是为了兼容C语言中的case设计的
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")		//旧代码 不要写 看懂就行
		fallthrough
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}

	// 跳出多重for循环
	// 通过bool值
	// goto+label实现跳出多层for循环

}
