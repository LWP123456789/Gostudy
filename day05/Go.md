### 切片

切片(Slice)是一个拥有相同类型元素的**可变长度**的序列;

它是**基于数组类型**做的一层封装；它非常灵活，支持自动扩容；

切片是一个引用类型，它的内部结构包含地址、长度、容量，切片一般用于快速地操作一块数据集合

#### 切片长度和容量的理解

切片指向了一个底层的数组

切片的长度就是它元素的个

切片的容量就是底层数组从切片的第一个元素到数组最后一个元素的数量

![image-20210820143357970](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20210820143357970.png)

```
切片是引用类型，都指向了底层的一个数组。
```

make()函数创造切片(指定长度)

```go
func main() {
	s1 := make([]int, 5, 10)
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

	s2 := make([]int, 0, 10)
	fmt.Printf("s2=%v len(s2)=%d cap(s2)=%d\n", s2, len(s2), cap(s2))
}
```

#### 切片的本质

切片就是一个框，框住了一块连续的内存

切片属于引用类型，真正的数据都是保存在底层数组里的。

#### 切片不能直接比较

切片之间是不能比较的，我们不能使用`==`操作符来判断两个切片是否含有全部相等元素。切片唯一合法的比较操作适合`nil`比较。一个`nil`值得切片并没有底层数组，一个`nil`值得切片的长度和容量都是0。但是我们不能说一个长度和容量都是0的切片一定是`nil`

```go
var s1 []int  		 //len(s1)=0;cap(s1)=0;s1==nil
s2 := []int{} 		 //len(s2)=0;cap(s2)=0;s1!=nil
s3 := make([]int,0)  //len(s3)=0;cap(s3)=0;s1!=nil
```

append()函数扩容

扩容策略：两倍      25%     超过两倍则到扩容量

Go扩容相关源代码

![Go_append_rule](D:\Image\Go_append_rule.jpg)

copy()函数拷贝切片

```go
copy(des,src)
```

#### Go语言中没有直接删除切片元素的方法

因此需要用append()方法将拆分的数组再合并从而达到删除的效果(相比其他语言确实麻烦...)

下面代码理解切片删除的底层数组发生的变化

```go
//实验
	x2 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 := x2[:]
	s2 = append(s2[:1], s2[2:]...)
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Println(x2)
```

```go
var a1 = [...]int{9, 156, 1, 5, 33}
	sort.Ints(a1[:]) //对切片进行排序
	fmt.Println(a1)
```

指针

Go语言中不存在指针操作，只需要记住两个符号：

1、&：取地址

2、*：根据地址取值

> 总结：取地址操作符`&`和取值操作符`*`是一对互补操作符，`&`取出地址，`*`根据地址去除地址指向的值



#### new函数申请一个内存地址(避免了空指针nil)



#### make和new的区别

1、make和new都是用来申请内存的

2、new很少用，一般用来给基本数据类型申请内存，`string`、 `int`，返回的是对应类型的指针(*string、*int)

3、make是用来给`slice`、`map`、 `chan`申请内存的，`make`函数返回的是对应的这三个类型本身

#### map

Go语言中提供的映射关系容器为`map`，其内部使用`散列表(hash)实现`

map是一种无序的基于`key-value`的数据结构，Go语言中的map是引用类型，必须初始化才能使用

##### !!!map遍历顺序是随机的

![image-20210820172847651](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20210820172847651.png)

删除操作：

```go
//删除
	delete(m1, "lep101")
	fmt.Println(m1)
```

练习题：

```go
// 统计一个字符串中每个单词出现的次数，比如：“how do you do”中how=1，do=2 you=1。
str := "how do you do"
m2 := make(map[string]int, 10)
str2 := strings.Split(str, " ")
/*fmt.Println(m2[str2[0]], m2[str2[1]], m2[str2[2]])*/
for _, v := range str2 {
   m2[v]++
}
fmt.Println(m2)
fmt.Printf("typeof str: %T typeof str2: %T", str, str2)
```

#### 函数

见代码... 一堆类型...

