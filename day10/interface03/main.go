package main

import "fmt"

// 接口的实现

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

func (c cat) move() {
	fmt.Println("走猫步...")
}

func (c cat) eat(food string) {
	fmt.Printf("猫吃%s...\n", food)
}

type chicken struct {
	feet int8
}

func (c chicken) move() {
	fmt.Println("鸡动!")
}

func (c chicken) eat(food string) {
	fmt.Printf("吃%s\n", food)
}

func main() {

	var a1 animal // 定义一个接口类型的变量

	bc := cat{ // 定义一个cat类型的变量bc
		name: "淘气",
		feet: 4,
	}

	a1 = bc
	a1.move()
	a1.eat("鱼")
	kfc := chicken{
		feet: 2,
	}
	a2 := kfc
	a2.move()
	a2.eat("1")
}

// 实现了接口变量能够存储不同的值.
