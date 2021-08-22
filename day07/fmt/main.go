package main

import "fmt"

// fmt

func main() {
	fmt.Print("lep101")
	fmt.Print("lep101.top")
	fmt.Println("----------------")
	fmt.Println("lep101")
	fmt.Println("lep101.top")

	var m1 = make(map[string]int, 1)
	m1["lep101"] = 100
	fmt.Printf("%v\n", m1)
	fmt.Printf("%#v\n", m1)

	num := 90
	fmt.Printf("%d%%\n", num) //在Printf里的格式化里用%做转义符

	fmt.Printf("%v\n", 100) // 通用%v万能，当不知道什么类型就用%v
	// 整数 -> 字符
	fmt.Printf("%q\n", 65)
	// 浮点数和复数
	fmt.Printf("%b\n", 3.14159265354697)
	// 字符串
	fmt.Printf("%q\n", "lep101")
	// 宽度标识符
	n := 12.34
	fmt.Printf("%f\n", n)
	fmt.Printf("%9f\n", n)
	fmt.Printf("%.2f\n", n)
	fmt.Printf("%9.2f\n", n)
	fmt.Printf("%9.f\n", n)

	// 获取用户输入
	var s string
	fmt.Scan(&s) // 具体用法有待研究
	fmt.Println("用户输入的内容是：", s)

	var (
		name  string
		age   int
		class string
	)
	fmt.Scanf("%s %d %s\n", &name, &age, &class)
	fmt.Println(name, age, class)

	fmt.Scanln(&name, &age, &class)
	fmt.Println(name, age, class)
}
