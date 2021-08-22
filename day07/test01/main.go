package main

import (
	"fmt"
)

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

// 主函数内要尽量清晰简洁
func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
	for i, j := range distribution {
		fmt.Printf("%s : %d\n", i, j)
	}
}

func dispatchCoin() int {
	//1、依次拿到每个人的名字
	for _, name := range users {
		//2、拿到一个人名根据分金币的规则去分金币
		for _, n := range name {
			switch n {
			case 'e', 'E': // 一个汉字和一个字母都属于字符
				distribution[name]++
			case 'i', 'I':
				distribution[name] += 2
			case 'o', 'O':
				distribution[name] += 3
			case 'u', 'U':
				distribution[name] += 4
			}
		}
		//2.1、每个人分的金币数应该保存到distribution中

		//2.2、还要记录下剩余的金币数
		coins -= distribution[name]
	}
	//3、整个第2步执行完就能得到最终每个人分的金币数和剩余金币数
	return coins
}

// 做题思路：先得到每个人的名字，
//在每个人的名字里遍历分情况判断，每判断完一个人的名字包含的金币后还要对剩下金币进行统计
