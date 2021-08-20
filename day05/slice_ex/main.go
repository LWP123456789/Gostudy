package main

import (
	"fmt"
	"sort"
)

// 切片练习题

func main() {
	var a = make([]int, 5, 10)
	fmt.Println(a)
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a, cap(a))

	var a1 = [...]int{9, 156, 1, 5, 33}
	sort.Ints(a1[:]) //对切片进行排序
	fmt.Println(a1)
}
