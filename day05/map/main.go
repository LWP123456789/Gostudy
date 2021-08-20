package main

import "fmt"

// map

func main() {
	m1 := make(map[string]int, 10) // 初始化时要估算好该map容量，避免在程序运行期间再动态扩容

	m1["乐培"] = 20
	m1["lep101"] = 35

	/*fmt.Println(m1)*/
	// 约定俗成用ok接受返回的布尔值
	/*value, ok := m1["娜扎"]
	if !ok {
		fmt.Println("查无此key")
	} else {
		fmt.Println(value)
	}*/

	//map的遍历
	/*for k, v := range m1 {
		fmt.Println(k, v)
	}
	// 只遍历key
	for v := range m1 {
		fmt.Println(v)
	}
	// 只遍历value
	for _, v := range m1 {
		fmt.Println(v)
	}*/
	//删除
	delete(m1, "lep101")
	fmt.Println(m1)
	//删除不存在的key delete函数误操作
}
