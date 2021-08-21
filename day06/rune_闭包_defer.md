## rune类型

rune是Go语言中一种特殊的数据类型,它是`int32`的别名,几乎在所有方面等同于`int32`,用于**区分字符值和整数值**

如果想要获得真实的字符串长度而不是其所占用字节数,有两种方法实现

方法一:

使用`unicode/utf-8`包中的`RuneCountInString`方法

```go
str := "hello 世界"
fmt.Println("RuneCountInString:", utf8.RuneCountInString(str))
```

方法二:

将字符串转换为rune类型的数组再计算长度

```go
str := "hello 世界"
fmt.Println("rune:", len([]rune(str)))
```

## defer

Go语言中的`defer`语句会将其后面跟随的语句进行延迟处理。在`defer`归属的函数即将返回时，将延迟处理的语句按`defer`定义的逆序进行执行，也就是说，先被`defer`的语句最后被执行，最后被`defer`的语句，最先被执行。

举个例子：

```go
func main() {
	fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")
}
```

输出结果：

```go
start
end
3
2
1
```

由于`defer`语句延迟调用的特性，所以`defer`语句能非常方便的处理资源释放问题。比如：资源清理、文件关闭、解锁及记录时间等。

```go
package main

import "fmt"

// 面试题

func calc(index string, a, b int) int {
   ret := a + b
   fmt.Println(index, a, b, ret)
   return ret
}

func main() {
   x := 1
   y := 2
   defer calc("AA", x, calc("A", x, y))
   x = 10
   defer calc("BB", x, calc("B", x, y))
   y = 20
}

// 1、 x := 1
// 2、 y := 2
// 3、 defer calc("AA",1,calc("A",1,2))
// 4、 calc("A",1,2) // => A 1 2 3
// 5、 defer calc("AA",1,3)
// 6、 x = 10
// 7、 defer calc("BB",10,calc("B",10,2))
// 8、 calc("B",10,2) // => B 10 2 12
// 9、 defer calc("BB",10,12)
// 10、y = 20
// 11、calc("BB",10,12) // => BB 10 12 22
// 11、calc("AA",3,4) // => AA 1 3 4
```

## 高阶函数

高阶函数分为函数作为参数和函数作为返回值两部分。

### 函数作为参数

函数可以作为参数：

```go
func add(x, y int) int {
	return x + y
}
func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}
func main() {
	ret2 := calc(10, 20, add)
	fmt.Println(ret2) //30
}
```

### 函数作为返回值

函数也可以作为返回值：

```go
func do(s string) (func(int, int) int, error) {
	switch s {
	case "+":
		return add, nil
	case "-":
		return sub, nil
	default:
		err := errors.New("无法识别的操作符")
		return nil, err
	}
}
```

## 匿名函数和闭包

### 匿名函数

函数当然还可以作为返回值，但是在Go语言中函数内部不能再像之前那样定义函数了，只能定义匿名函数。匿名函数就是没有函数名的函数，匿名函数的定义格式如下：

```go
func(参数)(返回值){
    函数体
}
```

匿名函数因为没有函数名，所以没办法像普通函数那样调用，所以匿名函数需要保存到某个变量或者作为立即执行函数:

```go
func main() {
	// 将匿名函数保存到变量
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	add(10, 20) // 通过变量调用匿名函数

	//自执行函数：匿名函数定义完加()直接执行
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)
}
```

匿名函数多用于实现回调函数和闭包。

### 闭包

闭包指的是一个函数和与其相关的引用环境组合而成的实体。简单来说，`闭包=函数+引用环境`。 首先我们来看一个例子：

```go
func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}
func main() {
	var f = adder()
	fmt.Println(f(10)) //10
	fmt.Println(f(20)) //30
	fmt.Println(f(30)) //60

	f1 := adder()
	fmt.Println(f1(40)) //40
	fmt.Println(f1(50)) //90
}
```

变量`f`是一个函数并且它引用了其外部作用域中的`x`变量，此时`f`就是一个闭包。 在`f`的生命周期内，变量`x`也一直有效。 闭包进阶示例1：

```go
func adder2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}
func main() {
	var f = adder2(10)
	fmt.Println(f(10)) //20
	fmt.Println(f(20)) //40
	fmt.Println(f(30)) //70

	f1 := adder2(20)
	fmt.Println(f1(40)) //60
	fmt.Println(f1(50)) //110
}
```

闭包进阶示例2：

```go
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func main() {
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test")) //test.jpg
	fmt.Println(txtFunc("test")) //test.txt
}
```

闭包进阶示例3：

```go
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	f1, f2 := calc(10)
	fmt.Println(f1(1), f2(2)) //11 9
	fmt.Println(f1(3), f2(4)) //12 8
	fmt.Println(f1(5), f2(6)) //13 7
}
```

闭包其实并不复杂，只要牢记`闭包=函数+引用环境`。

```go
// 闭包返回值是一个函数 在返回函数的外头的变量是不会在函数运行完后立即被销毁的
// 相当于开辟了一个内存存放参数，而这个参数值会被返回函数所影响

func adder(x int) func(int) int {
   return func(y int) int {
      x += y
      return x
   }
}

func main() {
   ret := adder(320)
   fmt.Println(ret(200))
   ret2 := ret(200)
   fmt.Println(ret2)
}
```

**闭包的好处：**

1.希望变量长期驻扎在内存当中（一般函数执行完毕，变量和参数会被销毁）

2.避免全局变量的污染