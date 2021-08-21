package main

import "fmt"

// defer

func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("lep101") //defer把它后面的语句延迟到函数即将返回的时候再执行
	fmt.Println("end")
}

// Go语言中函数的return不是原子操作，在底层分为两步来执行
// 第一步：返回值赋值
// 第二步：真正的RET返回
// 函数中如果存在defer，那么defer执行的时机是在第一步和第二步之间

func f1() int { //没有命名的返回值
	x := 5
	defer func() {
		x++ //修改的是x不是返回值
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x // 返回值 = y = x = 5
}

func f4() (x int) {
	defer func(x int) { //返回值
		x++ // 改变的是函数中的副本
	}(x) //传入的参数
	return 5 // 返回值 = x = 5
}

func main() {
	deferDemo()
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())

}
