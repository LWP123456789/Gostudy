package main

import "fmt"

// if条件判断
func main()  {
	/*age := 20

	if age > 18 {
		fmt.Println("澳门首家线上赌场开业啦！")
	}else {
		fmt.Println("该写暑假作业啦！")
	}

	// 多个判断条件
	if age > 35 {
		fmt.Println("人到中年")
	}else if age > 18 {
		fmt.Println("青年才俊")
	}else {
		fmt.Println("好好学习!")
	}*/

	//作用域
	//age变量此时只在if条件判断语句中生效 减小程序内存的占用
	if age := 19; age > 18 {
		fmt.Println(1)
	}else {
		fmt.Println(2)
	}
	/*fmt.Println(age)//此处找不到age*/

	//for 循环
	//Go语言中的所有循环类型均可以使用for关键字来完成
	//基本格式
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//变种1
	/*var i = 5
	for  ;i < 10; i++ {
		fmt.Println(i)
	}*/

	//变种2
/*	var i = 5
	for i < 10 {
		fmt.Println(i)
		i++
	}*/

	//无限循环 不要尝试
	/*for  {
		fmt.Println("123")
	}*/

	//for range(键值循环)
	s := "Hello乐培"
	for i,v := range s {
		fmt.Printf("%d %c\n",i,v)
	}
	//九九乘法表
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d ",i,j,i*j)
		}
		fmt.Println()
	}
 }
