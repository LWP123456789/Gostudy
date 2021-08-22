package main

import "fmt"

// 结构体

type Person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	// 声明一个Person类型的变量lep
	var lep Person
	// 通过字段赋值
	lep.name = "乐培"
	lep.age = 20
	lep.gender = "男"
	lep.hobby = []string{"篮球", "足球", "双色球"}
	fmt.Println(lep)
	// 访问变量lep的字段
	fmt.Println(lep.name)

	var lep101 Person
	lep101.name = "lwp"
	lep101.age = 20
	fmt.Printf("type:%T value:%v\n", lep101, lep101)

	// 匿名结构体:多用于临时场景 使用次数少 节省空间
	var s struct {
		x string
		y int
	}
	s.x = "乐乐乐乐培"
	s.y = 100
	fmt.Printf("type:%T value:%v\n", s, s)

}
