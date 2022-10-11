package modles

import (
	"encoding/csv"
	"fmt"
	"os"
	"project/call_the_roll/rollbook/utils"
	"strconv"
)

var AbsentCount int = 0
var dir string = "../output/"

//学生类
type student struct {
	id        int
	isPresent bool
}

//点名册类
type rollBook [90]student

//课程类
type course struct {
	name       string
	rollbook   rollBook
	badStudent []int //缺课多的学生
	//	numBadStudent	int
	count int //上课次数
}

//5门课程
type courseGroup [5]course

//初始化点名册
func rollBookDefault() (rollbook rollBook) {
	for i := 0; i < 90; i++ {
		rollbook[i].isPresent = true
		rollbook[i].id = i
	}
	return rollbook
}

//刷新点名册
func (rollbook *rollBook) flashBollbook() {
	for i := 0; i < 90; i++ {
		rollbook[i].isPresent = true
	}
}

//初始化课程类
func courseDefault(name string) (c course) {
	c.name = name
	c.count = 0
	c.rollbook = rollBookDefault()

	//rand.Seed(time.Now().Unix())
	//num := rand.Intn(4) + 5 //5~8

	num := utils.Random(4) + 5
	//fmt.Println(num)
	//	c.numBadStudent = num
	for i := 0; i < num; i++ { //随机num个id为badstudent

		//id := rand.Intn(90)
		id := utils.Random(90)
		c.badStudent = append(c.badStudent, id)
	}

	//fmt.Println(c.name, c.badStudent)
	return c
}

//模拟一次到课情况
func (c *course) actualPresent() {
	for _, id := range c.badStudent {
		if utils.EightyProbablity() { //缺席
			c.rollbook[id].isPresent = false
			AbsentCount++
		} else { //没缺席
			c.rollbook[id].isPresent = true
		}
	}

	//num := rand.Intn(3)
	num := utils.Random(4)
	for i := 0; i < num; i++ {
		//id := rand.Intn(90)
		id := utils.Random(90)
		c.rollbook[id].isPresent = false
		AbsentCount++
	}

}

//导出一次到课情况并重置roolbook
func (c *course) outputRollbookToCSV() {
	dsn := dir + c.name + "/" + c.name + strconv.Itoa(c.count) + ".csv"
	file, err := os.Create(dsn)
	if err != nil {
		fmt.Println("create file failed, err: ", err)
	}
	defer file.Close()

	w := csv.NewWriter(file)
	for _, student := range c.rollbook {
		err := w.Write([]string{
			strconv.Itoa(student.id),
			strconv.FormatBool(student.isPresent), //fmt.Sprintln(student.isPresent),
		})
		if err != nil {
			fmt.Println("write data failed, err: ", err)
			break
		}
	}
	w.Flush()

	c.rollbook.flashBollbook()
	c.count++
}

//初始化课程组
func CourseGroupDefault() (cgp courseGroup) {
	cgp[0] = courseDefault("A")
	cgp[1] = courseDefault("B")
	cgp[2] = courseDefault("C")
	cgp[3] = courseDefault("D")
	cgp[4] = courseDefault("E")
	return cgp
}

//模拟一组课的到课情况
func (cgp *courseGroup) ActualPresent() {
	for i := 0; i < 5; i++ {
		cgp[i].actualPresent()
		cgp[i].outputRollbookToCSV()
	}
}
