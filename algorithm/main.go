package main

import (
	"algorithm/modles"
)

func main() {
	//初始化课程组
	group := modles.DefaultCourseGroup()
	//抽点
	group.Exec()
}
