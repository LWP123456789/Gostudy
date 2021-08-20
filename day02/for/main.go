package main

import (
	"fmt"
)

func main()  {
	s := "Hello乐培大帅比"
	n := len(s)
	fmt.Println(n)

	/*for i := 0; i < len(s); i++  {
		fmt.Println(s[i])
		fmt.Printf("%c\n",s[i])	// %c:字符
	}*/
	/*for _, c:= range s {
		fmt.Printf("%c\n",c) // %c: 字符
	}*/

	// 字符串转化修改  字符串是一个不能修改的变量 只能通过转换去修改
	s2 := "白萝卜"	// => '白' '萝' '卜'
	s3 := []rune(s2) //把字符串强制转换成了一个rune切片
	s3[0] = '红'  //双引号表示字符串，单引号表示字符
	fmt.Println(string(s3)) //把rune切片强制转换成字符串

	c1 := "红"
	c2 := '红' //rune(int32)
	fmt.Printf("c1:%T c2:%T\n",c1, c2)
	c3 := "H"
	c4 := 'H'
	fmt.Printf("c3:%T c4:%T\n",c3, c4)

	//类型转换
	n1 := 10
	var f float64
	f = float64(n1)
	fmt.Println(f)


}


