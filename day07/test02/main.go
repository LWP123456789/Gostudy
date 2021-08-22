package main

import "fmt"

// Go语言坑之 for range
type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	/*fmt.Printf("%T\n", stus)*/
	/*fmt.Printf("%p\n", &stus)*/

	for _, stu := range stus {
		/*name := stu
		m[stu.name] = &name
		fmt.Printf("%p\n", &name)
		fmt.Printf("%p\n", &stu)*/
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}

// 1、对切片stus进行遍历，该切片指向的是一个结构体类型数组
// 2、m[小王子] = &stu
