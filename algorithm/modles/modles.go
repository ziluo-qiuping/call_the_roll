package modles

import (
	"algorithm/utils"
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const M int = 10 //每轮抽点人数
var dir string = "../output/"

//学生类
type student struct {
	id          int
	absentCount int    //缺课次数
	stringcode  string //出勤情况 1：上课 0：缺课
}

//班级类类
type class [90]student

//课程类
type course struct {
	name  string
	class class
	count int //上课次数
	sum   int //总点名
	valid int //有效次数
}

//课程组类
type courseGroup [5]course

//点名类
type record struct {
	id        int
	isPresent bool
}

//初始化班级学生信息
func (c *class) defauleClass() {
	for i := 0; i < 90; i++ {
		c[i].id = i
		c[i].stringcode = ""
		c[i].absentCount = 0
	}
}

//初始化课程
func defaultCourse(Name string) (c course) {
	c.name = Name
	c.class.defauleClass()
	return
}

//读取一次到课情况
func (c *course) inputRollbook() (rollbook []record) {
	dsn := dir + c.name + "/" + c.name + strconv.Itoa(c.count) + ".csv"
	file, err := os.Open(dsn)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 初始化一个 csv reader，并通过这个 reader 从 csv 文件读取数据
	reader := csv.NewReader(file)
	// 设置返回记录中每行数据期望的字段数，-1 表示返回所有字段
	reader.FieldsPerRecord = -1
	// 通过 readAll 方法返回 csv 文件中的所有内容
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	for _, item := range records {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		r := record{int(id), true}
		bool := item[1]
		//fmt.Println(bool)
		if bool == "false" {
			r.isPresent = false
		}
		rollbook = append(rollbook, r)
	}
	return
}

//进行一次算法
func (c *course) algorithm() {
	rb := c.inputRollbook() //读入一次课程记录
	if c.count < 1 {        //第一次全读取 
		for _, v := range rb {
			if v.isPresent == false { //没到
				c.class[v.id].absentCount++
				//c.class[v.id].stringcode += "0"
				c.valid++
			} //} else {  //有到
			//	c.class[v.id].stringcode += "1"
			//}
		}
		c.output(0, []int{}, rb)
		c.sum += 90
	} else { //第一次之后 先找出缺课最多的前n个学生 在随机找M-n个学生  目前M设置为10，放弃了随机查找学生
		n := 10 //n课根据上课次数修改
		c.findAbsent(rb, n, M)
	}
	c.count++
}

//第一次之后 先找出缺课最多的前n个学生 在随机找m-n个学生
func (c *course) findAbsent(rb []record, n int, m int) {
	ids := findAbsentMax(c.class, n)
	c.output(0, ids, rb)
	for _, id := range ids {
		if rb[id].isPresent == false {
			c.class[id].absentCount++
			c.valid++
		}
	}
	for i := 0; i < m-n; i++ {
		id := utils.Random(90)
		if utils.IsIN(id, ids) { //判断id是否已经被检查过了
			i = i - 1
			continue
		} else {
			ids = append(ids, id)
			if rb[id].isPresent == false {
				c.class[id].absentCount++
				c.valid++
			}
		}
	}
	c.output(1, ids, rb)
	c.sum += m
}

//找出前n个缺课最多的学生
func findAbsentMax(class class, n int) (ids []int) {

	var top []student = class[:]
	sort.Slice(top, func(i, j int) bool {
		return top[i].absentCount > top[j].absentCount
	})

	// fmt.Println("top")
	// fmt.Println(top)

	for i := 0; i < n; i++ {
		ids = append(ids, top[i].id)
	}

	return ids
}

//初始化课程组
func DefaultCourseGroup() (group courseGroup) {
	group[0] = defaultCourse("A")
	group[1] = defaultCourse("B")
	group[2] = defaultCourse("C")
	group[3] = defaultCourse("D")
	group[4] = defaultCourse("E")
	return
}

//组执行一次算法
func (cg courseGroup) Exec() {
	var total_sum, total_valid int = 0, 0
	for _, c := range cg {
		c.exec()
		total_sum += c.sum
		total_valid += c.valid
	}
	fmt.Printf("total sum: %d total valid: %d E: %f \n", total_sum, total_valid, float64(total_valid)/float64(total_sum))
}

// 进行20次算法
func (c *course) exec() {
	for i := 1; i <= 20; i++ {
		c.algorithm()
		
			fmt.Printf("course: %s count: %d sum: %d valid: %d E: %f \n", c.name, c.count, c.sum, c.valid, float64(c.valid)/float64(c.sum))
		
	}
}

//输出抽点方案到list目录ixa
func (c *course) output(pattern int, ids []int, rb []record) {
	dsn := "../list/" + c.name + "/" + c.name + strconv.Itoa(c.count) + ".csv"
	file, err := os.Create(dsn)
	if err != nil {
		fmt.Println("create file failed, err: ", err)
	}
	defer file.Close()

	w := csv.NewWriter(file)
	if pattern == 0 { //前三次全抽
		for _, student := range rb {
			err := w.Write([]string{
				strconv.Itoa(student.id),
				strconv.FormatBool(student.isPresent), //fmt.Sprintln(student.isPresent),
			})
			if err != nil {
				fmt.Println("write data failed, err: ", err)
				break
			}
		}
	} else { //之后抽ids
		for _, id := range ids {
			err := w.Write([]string{
				strconv.Itoa(rb[id].id),
				strconv.FormatBool(rb[id].isPresent),
			})
			if err != nil {
				fmt.Println("write data failed, err: ", err)
				break
			}
		}
	}
	w.Flush()
}
