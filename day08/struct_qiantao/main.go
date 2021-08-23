package main

import "fmt"

// 嵌套结构体

type address struct {
	province string
	city     string
}

/*type workPlace struct {
	province string
	city     string
}*/

type person struct {
	name    string
	age     int
	address // 匿名嵌套结构体  防止冲突时需访问好成员 例如workPlace和address，老老实实写全
	// address:address
}

/*
type company struct {
	name string
	addr address
}*/

func main() {
	p1 := person{
		name: "lwp",
		age:  20,
		address: address{
			province: "广东",
			city:     "佛山", // 语法糖：结构体匿名嵌套可以节省代码中成员的访问
		},
	}
	fmt.Println(p1)
	fmt.Println(p1.name, p1.city)
}
