package main

import (
	"fmt"
	"strings"
)

// Go语言中推荐使用驼峰式命名:
// var studentName string
//变量 先声明 后使用
/*
var name string //声明一个保存字符串类型数据的s1类型
var age int
var isOk bool*/

//批量声明变量
var (
	name string // ""
	age int     // 0
	isOk bool   // false
)
//批量声明常量
const (
	STATUSOK = 200
	NOTFOUND = 404
)
// 如果批量声明常量时，如果某一行声明后没有赋值，默认就和上一行一致
const (
	n1 = 100
	n2
	n3
)

// iota：实现枚举
const (
	a1 = iota//0 当索引用
	a2 //1
	a3
)

const (
	b1 = iota // 0
	b2		  // 1
	_         // 2
	b3        // 3
)

//插队
const (
	c1 = iota
	c2 = 100
	c3 =iota
	c4
)

// 多个常量声明在一行
const (
	d1,d2 = iota + 1, iota + 2
	d3,d4 = iota + 1, iota + 2
)

// 定义数量级
const (
	_ = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main()  {
	name = "乐培"
	age = 20
	isOk = true
	// Go 语言中非全局变量声明后必须使用
	fmt.Print(isOk) //在终端中输出要打印的内容
	fmt.Printf("\nname:%s",name) // %s:占位值 使用name这个变量的值去替换占位符
	fmt.Println(age) //打印完指定的内容之后会在后面加一个换行符
	fmt.Println("n1:",n1)
	fmt.Println("n2:",n2)
	fmt.Println("n3:",n3)

	fmt.Println("a1:",a1)
	fmt.Println("a2:",a2)
	fmt.Println("a3:",a3)

	fmt.Println("b1:",b1)
	fmt.Println("b2:",b2)
	fmt.Println("b3:",b3)

	fmt.Println("c1:",c1)
	fmt.Println("c2:",c2)
	fmt.Println("c3:",c3)
	fmt.Println("c4:",c4)

	fmt.Println("d1:",d1)
	fmt.Println("d2:",d2)
	fmt.Println("d3:",d3)
	fmt.Println("d4:",d4)
	//声明变量同时赋值
	var s1 string = "乐乐乐乐培"
	fmt.Println(s1)
	//类型推导(根据值判断变量是什么类型)
	var s2 = "20"
	fmt.Println(s2)
	//简短变量声明 := 只能在函数里面用
	s3 := "哈哈哈"
	fmt.Println(s3)
	//同一个作用域({})中不能重复声明同名的变量
	//匿名变量是一个特殊的变量 :_

	i1 := 101
	fmt.Printf("%d\n", i1)
	fmt.Printf("%b\n", i1) // 把十进制数转换成二进制
	fmt.Printf("%o\n", i1) // 把十进制数转换成八进制
	fmt.Printf("%x\n", i1) // 把十进制数转换成十六进制
	// 八进制
	i2 := 077
	fmt.Printf("%d\n", i2)
	// 十六进制
	i3 := 0x1234567
	fmt.Printf("%d\n", i3)
	// 查看变量的类型
	fmt.Printf("%T\n",i3)

	// 声明int8类型的变量
	i4 := int8(9) // 明确指定int8类型，否则就是默认认为int类型
	fmt.Printf("%T\n",i4)
	/*
	math.MaxFloat32 //float32最大值*/
	f1 := 1.23456
	fmt.Printf("%T\n",f1) // 默认Go语言中的小数都是float64类型
	f2 := float32(1.23456)
	fmt.Printf("%T\n",f2) //显示声明float32类型
	/*f1 = f2 // float32类型的值不能直接赋值给float64类型*/

	//复数
	var z1 complex64
	z1 = 1 + 2i
	var z2 complex128
	z2 = 2 + 3i
	fmt.Println(z1)
	fmt.Println(z2)

	//布尔值
	/*
		1、布尔类型变量的默认值为false
		2、Go语言中不允许将整形强制转换为布尔型
		3、布尔型无法参与数值运算，也无法与其他类型进行转换
	*/
	b1 := true
	var  b2 bool //默认是false
	fmt.Printf("%T\n",b1)
	fmt.Printf("%T value:%d\n",b2,b2)

	//fmt占位符
	//查看类型 %T %v(查看变量的值,万能) %b %d %o %x %s
		var s = "Hello 乐培"
		fmt.Printf("%s\n",s)
		fmt.Printf("%v\n",s)
		fmt.Printf("%#v\n",s)//#显示类型，比如字符串会给值显示一个双引号

	//字符串
	//Go语言中字符串是用双引号包裹的！！！
	//Go语言中单引号包裹的是字符！！！
	/*s := "Hello 乐培"*/
	// 单独的字母，汉字，符号表示一个字符
	/*c1 := 'h'*/
	// \ 本来是具有特殊含义的    转义字符
	path := "\"D:\\Go_WorkSpace\\src\\github.com\\LWP123456789\""
	fmt.Printf("%v\n", path)

	// 多行的字符串
	s4 := `
		世情薄
		人情恶
		雨送黄昏花易落
	`
	fmt.Println(s4)
	path2 := `"D:\Go_WorkSpace\src\github.com\LWP123456789"`
	fmt.Println(path2)

	//字符串相关操作 常用操作
	name := "乐培"
	adj := "大帅比"
	//拼凑
	ss := name + adj
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s%s",name,adj)	// 把俩字符串拼接在一起形成一个新的字符串
	fmt.Printf("%s%s",name,adj)   // 返回到终端
	fmt.Println(ss1)
	// 分隔
	ret := strings.Split(path2,"\\")
	fmt.Println(ret)
	// 包含
	fmt.Println(strings.Contains(ss,"乐培")) //某字符串包含某字符则返回true 否则返回false
	// 前缀  是不是以 某字符 开头
	fmt.Println(strings.HasPrefix(ss,"乐培"))
	// 后缀  是不是以 某字符 结尾
	fmt.Println(strings.HasSuffix(ss,"比"))
	// Index 判断并返回字串的位置
	s5 := "abcden"
	fmt.Println(strings.Index(s5,"c"))
	// 拼接
	fmt.Println(strings.Join(ret,"+"))
}

