package main

import (
	"fmt"
	"unicode"
)

func main()  {

	var count int
	s1 := "Hello 乐培大帅比"

	for _, c := range s1 {
		if unicode.Is(unicode.Han,c) {
			count++
		}
	}

	fmt.Println(count)
}
//unicode.Is(unicode.Han, c)
//1. unicode.Han包含了所有的汉字
//2.unicode.Is(unicode.Han, c) 这是判断 c是否在unicode.Han中
//3.如果在unicode.Is(unicode.Han, c) 这个函数返回true 否则返回false