package main

import (
	"fmt"
	"strings"
)

// map和slice组合

func main() {
	// 元素类型为map的切片
	var s1 = make([]map[int]string, 1, 10)
	//没有对内部的map做初始化
	s1[0] = make(map[int]string, 1)
	s1[0][10] = "A"
	fmt.Println(s1)

	// 值为切片类型的map
	var m1 = make(map[string][]int, 10)
	m1["北京"] = []int{10, 20, 30}
	fmt.Println(m1)

	// 统计一个字符串中每个单词出现的次数，比如：“how do you do”中how=1，do=2 you=1。
	str := "how do you do"
	m2 := make(map[string]int, 10)
	str2 := strings.Split(str, " ")
	/*fmt.Println(m2[str2[0]], m2[str2[1]], m2[str2[2]])*/
	for _, v := range str2 {
		m2[v]++
	}
	fmt.Println(m2)
	fmt.Printf("typeof str: %T typeof str2: %T", str, str2)
}
