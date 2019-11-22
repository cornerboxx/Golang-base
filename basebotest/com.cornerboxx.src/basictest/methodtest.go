package main

import "fmt"

func main() {
	//要执行同包下的其他方法，使用命令"go build  basetest.go methodtest.go"
	//注意这两个文件中只允许一个main方法，否则报错
	//logic1()

	//值传递方法
	int1, int2 := 1, 2
	changeNumValue(int1, int2)
	fmt.Println(int1, int2)
	//引用传递方法
	changeNumValueReal(&int1, &int2)
	fmt.Println(int1, int2)

}

func changeNumValue(int1, int2 int) {
	temp := int2
	int2 = int1
	int1 = temp
}

func changeNumValueReal(int1, int2 *int) {
	fmt.Println(int1, int2)
	temp := *int2
	*int2 = *int1
	*int1 = temp
}
