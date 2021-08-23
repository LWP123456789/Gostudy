package main

import "fmt"

// 方法

type person struct {
	name string
	age  int
}

// Go语言中如果标识符首字母是大写的，就表示对外部可见(暴露的，公有的)
// type Dog struct
type dog struct {
	name string
}

// 构造函数
func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

func newDog(name string) dog {
	return dog{
		name: name,
	}
}

// 方法是作用于特定类型的函数
// 接收者表示的是调用该方法的具体类型变量，多用类型名首字母小写表示
func (d dog) wang() {
	fmt.Printf("%s:汪汪汪~", d.name)
}

// 使用值接收者
func (p person) guonian() {
	p.age++
}

// 使用指针接收者  ：可以改变接收者中的值
func (p *person) zhenguonian() {
	p.age++
}

func (p *person) dream() {
	fmt.Println("不上班也能挣钱!")
}

// 给自定义类型加方法
// 不能给别的包里面的类型添加方法，只能给自己的包里的类型添加方法
type myInt int

func (m myInt) hello() {
	fmt.Println("我是一个", m)
}

func main() {
	d1 := newDog("lep")
	d1.wang()
	p1 := newPerson("lep", 20)
	fmt.Println(p1.age)
	p1.guonian()
	fmt.Println(p1.age)
	p1.zhenguonian()
	fmt.Println(p1.age)

	// 声明一个int32类型的变量，它的值是10
	// 方法1：var x int32 = 10
	// 方法2: var x int 32  x = 10
	// 方法3:
	var x = int32(10)
	// 方法4:
	// x := int32(10)
	fmt.Println(x)
	m := myInt(100) //强制类型转换
	m.hello()
}
