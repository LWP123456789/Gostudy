package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Go程序设计语言 练习1.1

func main() {
	for i, j := range os.Args {
		fmt.Printf("%v %v", i, j)
	}
	fmt.Println(strings.Join(os.Args[1:], " ")) // 减少数据处理的代价
	fmt.Println(os.Args[0])

	// 测试strings.join 与 普通追加操作之间 性能上的差别
	s := []string{"lep ", "is ", "a ", "very ", "handsome ", "man"}
	/*fmt.Println(s)*/
	temp := ""
	fmt.Println(time.Now())
	for _, j := range s {
		temp += j
	}
	fmt.Println(time.Now(), temp)
	fmt.Println(strings.Join(s, ""), time.Now())
}
