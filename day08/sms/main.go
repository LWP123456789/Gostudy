package main

import (
	"fmt"
	"os"
)

/*
	函数版学生管理系统
	写一个系统能够查看\新增学生\删除学生
*/

var (
	allStudent map[int64]*student // 变量
)

type student struct {
	id   int64
	name string
}

func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

func showAllStudents() {
	// 把所有的学生都打印出来
	for k, v := range allStudent {
		fmt.Printf("学号:%d 姓名:%s\n", k, v.name)
	}
}

func addStudent() {
	// 向allStudent中添加一个新的学生
	// 1、创建一个新学生
	var (
		id   int64
		name string
	)
	fmt.Print("请输入学生学号:")
	fmt.Scanln(&id)
	fmt.Print("请输入学生姓名:")
	fmt.Scanln(&name)
	newStu := newStudent(id, name)
	allStudent[id] = newStu

}

func deleteStudent() {
	//1、请用户输入要删除的学生的序号
	var (
		deleteId int64
	)
	fmt.Print("请输入学生学号:")
	fmt.Scanln(&deleteId)
	//2、去allStudent这个map中根据学号删除对应的键值对
	delete(allStudent, deleteId)
}

func main() {
	allStudent = make(map[int64]*student, 48) // 初始化(开辟内存空间)
	for {
		// 1、打印菜单
		fmt.Println("Welcome to lepのSys")
		fmt.Println(`
		1、查看所有学生
		2、新增学生
		3、删除学生
		4、退出
	`)
		fmt.Print("请输入你想要的操作:")
		// 2、等待用户选择要做什么
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了%d这个选项!\n", choice)
		// 3、执行对应函数
		switch choice {
		case 1:
			showAllStudents()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1) //退出
		default:
			fmt.Println("输入错误")
		}
	}
}
