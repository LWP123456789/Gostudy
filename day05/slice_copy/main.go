package main

import "fmt"

// copy

func main() {
	a1 := []int{1, 3, 5}
	a2 := a1
	/*var a3 []int //nil*/
	var a3 = make([]int, 3, 3)
	copy(a3, a1)
	fmt.Println(a1, a2, a3)
	a1[0] = 100
	fmt.Println(a1, a2, a3)

	// 将a1中的索引为1的3这个元素删掉
	a1 = append(a1[:1], a1[2:]...) // 只有一个元素也要拆分
	fmt.Println(a1)
	fmt.Println(cap(a1))

	x1 := [...]int{1, 3, 5} //数组
	s1 := x1[:]             //切片
	fmt.Println(s1, len(s1), cap(s1))
	s1 = append(s1[:1], s1[2:]...) //删除下标为1的元素 修改了底层数组！
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(x1, len(x1))
	//实验
	x2 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 := x2[:]
	s2 = append(s2[:1], s2[2:]...)
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Println(x2)
}
