# call_the_roll
2022秋软工实践 第二次结对编程作业

## 环境搭建

golang(version >= 1.17) （可选）

## 运行代码

#### 1.直接运行可执行文件（不用安装golang）

**generate.exe**:数据生成器，执行后将生成的数据以csv文件格式保存到output目录下，并输出总缺席次数

**main.exe**：算法执行，输出每门课每5次课的点名数，有效数，E，并在最后汇总5门课的总数据， 并生成抽点方案保存到list目录下

**2.运行代码（已安装golang）**

依次进入roolbook和algorithm目录下的main.go文件，运行`go mod tidy`，及`go run main.go`，效果同上

