package main

import "fmt"

func main()  {
	var (
		a = 5
		b = 2
	)

	// 算术运算符
	// + - * / %
	a++ // 单独的语句，不能放在=的右边赋值 ==> a = a + 1
	b-- // 单独的语句，不能放在=的右边赋值 ==> b = b - 1

	// 关系运算符
	fmt.Println(a == b) // Go语言是强类型，相同类型的变量才能比较
	fmt.Println(a != b) // 不等于
	fmt.Println(a >= b) //大于等于
	fmt.Println(a <= b) //小于等于
	fmt.Println(a < b) //小于
	//强/弱类型是指类型检查的严格程度。语言有无类型、弱类型、强类型3中。
	//
	//　　无类型的不检查，甚至不区分指令和数据。
	//
	//　　弱类型的检查很弱，仅能严格地区分指令和数据。
	//
	//　　强类型则严格在编译期进行检查。
	//
	//　　弱类型允许将一块内存看做多种类型。比如直接将整型变量与字符变量相加。
	//	  C和C++是静态语言，也是弱类型语言，Perl和PHP是动态语言，但也是弱类型语言。
	//	  强类型语言在没有强制类型转化前，不允许两种不同的类型变量相互操作。
	//	  Java，C#和Python都是强类型语言。

	// 逻辑运算符 && || !
	// 如果年龄大于18岁并且年龄小于60岁
	age := 22
	if age > 18 && age < 60 {
		fmt.Println("学习上班！")
	}else {
		fmt.Println("不用上班！")
	}
	// 如果年龄小于18岁 或者 年龄大于60岁
	if age < 18 || age > 60 {
		fmt.Println("不用上班!")
	}else {
		fmt.Println("学习上班!")
	}
	isMarried := false
	fmt.Println(isMarried)
	fmt.Println(!isMarried)

	// 位运算符 : 针对二进制数

	// &:按位与
	fmt.Println(5 & 2)
	// |:按位或
	fmt.Println(5 | 2)
	// ^:按位异或 (不一样则为 1)
	fmt.Println(5 ^ 2)
	// <<:将二进制左移指定位数
	fmt.Println(5 << 1)  // 1010 => 10
	fmt.Println(1 << 10) // 10000000000 => 1024
	// >>: 将二进制位右移指定的位数
	fmt.Println(5 >> 1)  //10 => 2
	fmt.Println(5 >> 2)

	// 赋值运算符
	var x = 10
	x += 1
	x -= 1
	x *= 1
	x /= 1
	x %= 1

	x <<= 2
	x &= 2
	x |= 3
	x ^= 4
	x >>= 2

	fmt.Println(x)
}
