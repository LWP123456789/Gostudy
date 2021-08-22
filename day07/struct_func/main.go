package main

import "fmt"

// 构造函数

type person struct {
	name string
	age  int
}

// 构造函数
// 返回的是结构体还是结构体指针 结构体内字段小返回值类型，多久返回指针类型，减小内存消耗
// 需要使用结构体多时使用构造函数，减小内存开销,减少代码
func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

func main() {
	p1 := newPerson("乐培", 20)
	p2 := newPerson("lep", 18)
	fmt.Println(p1, p2)
}
