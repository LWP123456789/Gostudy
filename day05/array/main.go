package main

import (
	"fmt"
)

// 数组

// 存放元素的容器
// 必须指定存放的元素的类型和容量(长度)
// 数组的长度是数组类型的一部分
func main() {
	var a1 [3]bool // [true false true]
	var a2 [4]bool // [true true false false]

	fmt.Printf("a1:%T a2:%T", a1, a2)

	// 数组的初始化
	fmt.Println(a1, a2) // 如果不初始化，默认元素都是零值(布尔值:false,整型浮点型0,字符串"")
	// 1、初始化方式1
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)
	// 2、初始化方式2
	// 根据初始值自动推断数组的长度是多少
	a10 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a10)
	// 3、初始化方式3:根据索引来初始化
	a3 := [5]int{0: 1, 4: 2}
	fmt.Println(a3)

	// 数组的遍历
	citys := [...]string{"北京", "上海", "深圳"}
	// 1、根据索引遍历
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}
	// 2、for range遍历
	for i, v := range citys {
		fmt.Println(i, v)
	}

	// 多维数组
	// [[1 2] [3 4] [5 6]]
	var a11 [3][2]int
	a11 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	} //Go语言中数组用空格来分割
	fmt.Println(a11)

	/*var a4 [4]int = [4]int{1,1,1,1}
	fmt.Println(a4)*/

	// 多维数组的遍历
	for _, v1 := range a11 {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	// 数组是值类型
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1, b2)

	// 求数组所有元素的和
	sum := 0
	for _, v := range a10 {
		sum += v
	}
	fmt.Println(sum)
	// 找出数组中和为指定值的两个元素的下标,比如从数组{1,3,5,7,8}中找出和为8的两个元素的下标
	// 分别为(0,3)和(1,2)
	a5 := [...]int{1, 3, 5, 7, 8}
	for i := 0; i < len(a5); i++ {
		for j := i + 1; j < len(a5); j++ {
			if a5[i]+a5[j] == 8 {
				fmt.Printf("(%d,%d)", i, j)
			}
		}
	}

}
