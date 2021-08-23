package main

//?反序列化传不了值
import (
	"encoding/json"
	"fmt"
)

// 结构体与json

// 1、序列化：把Go语言中的结构体变量 --> json格式的字符串
// 2、反序列化：json格式的字符串 --> Go语言中能够识别的结构体变量

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p1 := person{
		Name: "lep",
		Age:  20,
	}

	//序列化
	b, err := json.Marshal(p1) //json包要拿结构体中的成员需要大写或加上反引号 解析后用新的name取代旧的
	if err != nil {
		fmt.Printf("marshal failed,err%v", err)
		return
	}
	fmt.Printf("%#v\n", string(b))
	// 反序列化
	str := `{"name":"lep101","age",18}`
	var p2 person
	err = json.Unmarshal([]byte(str), &p2) // 传指针是为了在函数内部修改p2的值
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", p2)

}
