package main

import "fmt"

// 有一个物件：
// 1、它保存了一些数据		---> 结构体的字段
// 2、它有三个功能			---> 结构体的方法

type student struct {
	id   int64
	name string
}

// 造一个学生的管理者
type studentMgr struct {
	allStudent map[string]*student
}

/*var (
	allStudent map[string]student // 全局变量 存放学生数据
)
*/

// 查看学生
func (s studentMgr) showStudents() {
	// 从s.allstudent这个map中把所有的学生逐个拎出来
	for _, stu := range s.allStudent {
		fmt.Printf("学号:%d 姓名:%s\n", stu.id, stu.name)
	}
}

// 增加学生
func (s studentMgr) addStudent() {
	// 1、根据用户输入的内容创建一个新的学生
	var (
		stuId   int64
		stuName string
	)
	// 获取用户输入
	fmt.Print("请输入学号：")
	fmt.Scanln(&stuId)
	fmt.Print("请输入姓名：")
	fmt.Scanln(&stuName)
	// 根据用户输入创建结构体对象
	newStu := student{
		id:   stuId,
		name: stuName,
	}
	// 2、把新的学生放到s.allStudent这个map中
	s.allStudent[newStu.name] = &newStu
}

// 修改学生
func (s studentMgr) editStudent() {
	// 获取用户输入的姓名
	var stuName string
	fmt.Print("请输入姓名：")
	fmt.Scanln(&stuName)
	// 展示该学号对应的学生信息,如果没有提示查无此人
	stuObj, ok := s.allStudent[stuName]
	if !ok {
		fmt.Println("查无此人!")
	} else {
		fmt.Printf("你要修改的学生信息如下:学生:%d 姓名:%s\n", stuObj.id, stuObj.name)
		// 请输入修改后的学生ID
		fmt.Print("请输入学生的新ID:")
		var newId int64
		fmt.Scanln(&newId)
		// 更新学生的ID
		s.allStudent[stuName].id = newId
		/*s.allStudent[stuName] = stuObj // 更新map中的学生*/
	}
}

// 删除学生
func (s studentMgr) deleteStudent() {
	// 1.请用户输入要删除的学生名字
	var stuName string
	fmt.Print("请输入要删除学生的姓名：")
	fmt.Scanln(&stuName)
	// 2、去map找有没有这个学生，如果没有打印个查无此人的提示
	_, ok := s.allStudent[stuName]
	if !ok {
		fmt.Println("查无此人")
		return
	}
	// 3、有的话就删除,如何从map中删除键值对
	delete(s.allStudent, stuName)
	fmt.Println("删除成功!")
}
