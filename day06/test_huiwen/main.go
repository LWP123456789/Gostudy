package main

import "fmt"

// 回文判断

func main() {
	ss := "上海自来水来自海上"
	// 把字符串中的字符强制转换成[]rune类型
	r := []rune(ss)
	for i := range r {
		if r[i] != r[len(r)-1-i] {
			fmt.Println("不是回文")
			return
		}
	}
	fmt.Println("是回文")
}
