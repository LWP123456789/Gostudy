package main

import "fmt"

// 接口示例2
// 不管是什么牌子的车，都能跑

// 定义一个car接口类型
// 不管是什么结构体 只要有run方法都能是car类型
type car interface {
	run()
}

type Ferrari struct {
	brand string
}

func (f Ferrari) run() {
	fmt.Printf("%s速度70迈~\n", f.brand)
}

type Porsche struct {
	brand string
}

func (b Porsche) run() {
	fmt.Printf("%s速度700迈~", b.brand)
}

func drive(c car) {
	c.run()
}

func main() {
	var f1 = Ferrari{
		brand: "法拉利",
	}
	var b1 = Porsche{
		brand: "保时捷",
	}
	drive(f1)
	drive(b1)
}
