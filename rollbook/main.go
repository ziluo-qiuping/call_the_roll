package main

import (
	"fmt"
	"rollbook/modles"
)

func main() {
	//初始化5门课程
	cgp := modles.CourseGroupDefault()
	for i := 0; i < 20; i++ { //每门课程模拟20次到课情况
		cgp.ActualPresent()
	}
	//输出总缺席人数
	fmt.Printf("absent count: %d\n", modles.AbsentCount)
}
