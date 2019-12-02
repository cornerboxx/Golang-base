package main

import (
	"fmt"
	"strconv"
)

//定义一个函数类型
type fu func()
type fu1 func() int
type intType int

func main() {
	//要执行同包下的其他方法，使用命令"go build  basetest.go methodtest.go"
	//注意：这两个文件中只允许一个main方法，否则报错
	//logic1()

	//值传递方法
	int1, int2 := 1, 2
	changeNumValue(int1, int2)
	fmt.Println(int1, int2)
	//引用传递方法
	changeNumValueReal(&int1, &int2)
	fmt.Println(int1, int2)
	//函数作为参数
	body(func() {
	})
	body(callback)
	//闭包用法
	ff := closePak()
	fmt.Println(ff()())

	//函数方法，为一个传入的指针增加方法
	//注意：对int,string,func类型添加方法无效，推测是对系统提供的类型失效
	var boss Boss
	boss.smoking()
	boss.talk()
	//fu := func() {}
	//fu.talk()//报错
	//intType := 1
	//intType.talk()//报错

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

func body(f fu) {
	fmt.Println("这里是body")
	f()
}

func callback() {
	fmt.Println("这里是回调")
}

var iii int = 1

func closePak() func() fu1 {
	iii := 2
	str := strconv.Itoa(iii)
	fmt.Println(str)

	fun := func() int {
		var iii int = 3
		//就近原则，修改的是上一行的iii
		iii = 3
		fmt.Println(str)
		return iii
	}

	return func() fu1 {
		return fun
	}
}

type Boss struct {
	money float64
}

func (b Boss) fire() {
	fmt.Println(" boss get fire ")
}

func (b Boss) smoking() {
	b.fire()
	fmt.Println(" boss is smoking ")
}

func (b *Boss) talk() {
	fmt.Println(" boss is talking ")
}

//以下是几个失效的例子
func (b fu) talk() {
	fmt.Println(" fu can`t talk ")
}

func (b intType) talk() {
	fmt.Println(" intType can`t talk ")
}
