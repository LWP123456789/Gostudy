package main

import "fmt"

// 结构体是值类型

type person struct {
	name   string
	gender string
	age    int
}

// go语言中函数参数永远是拷贝
func f(x person) {
	x.gender = "女" // 修改的是副本的gender
}

func f2(x *person) {
	/*(*x).gender = "女" // 根据内存地址找到哪个原变量，修改的就是原来的变量*/
	x.gender = "女" // 语法糖，自动根据指针找对应的变量
}

func main() {
	var p person
	p.name = "乐培"
	p.gender = "男"
	p.age = 20
	f(p)
	fmt.Println(p.gender) // 男
	f2(&p)                // 传的是内存地址
	fmt.Println(p.gender) // 女
	// 结构体指针1
	var p2 = new(person)
	p2.name = "lep"
	fmt.Printf("%T\n", p2)
	fmt.Printf("%p\n", p2) // p2保存的值就是一个内存地址
	// 结构体指针2
	// 2.1、key-value初始化
	var p3 = person{
		name:   "元帅",
		gender: "男",
	}
	fmt.Printf("%v\n", p3)
	// 2.2、使用值列表的形式初始化,值的顺序要和结构体定义时字段的顺序一致
	p4 := person{
		"小乐培",
		"男",
		20,
	}
	fmt.Printf("%v\n", p4)
	fmt.Println(p4)
}
