package main

import "fmt"

// 递归:函数自己调用自己
// 递归适合处理那种问题相同\问题的规模越来越小的场景
// 递归一定要有一个明确的退出条件

// 计算n的阶乘
func f1(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return n * f1(n-1)
}

// 上台阶的面试题
// n个台阶，一次可以走1步，也可以走2步，有多少种走法
func f2(n uint64) uint64 {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else {
		return f2(n-1) + f2(n-2) // 剩一个台阶 和 剩两个台阶
	}
}

func main() {
	/*ret := f1(7)
	fmt.Println(ret)*/
	ret := f2(3)
	fmt.Println(ret)
}
