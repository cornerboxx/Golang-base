package main

import (
	"fmt"
	"unsafe"
)

//基本数据类型
var bool1 bool
var bool2, bool3 bool = true, false
var int1 int
var uint1 uint
var float321 float32
var float641 float64
var string1 string

//复杂数据类型
var int_ *int
var intArray []int
var mapLength map[string]int
var intChan chan int
var intFunc func(string) int
var errorRealize error // error 是接口

var (
	int2 int
	int3 int
)

//错误的声明方式
//var bool4 := true
//bool5 := true,false

//声明常量 常量不允许不赋值
const const1 int = 10
const const2 int = 5

//枚举 常量表达式中，函数必须是内置函数
const (
	enum1 = "abc"
	enum2 = len(enum1)
	enum3 = unsafe.Sizeof(enum2)
	//enum3 = newObj()
)

//iota 随行数增加，默认从0开始
//未赋值的默认取前一行的表达式
const (
	enum4 = 1 + iota
	enum5
)

//算术运算符（+ - * / % ++ --）
//关系运算符（ ==  != >= <= > < ）
//关系运算符（ ==  != >= <= > < ）
//逻辑运算符（ &&  || ! ）
//上述用法没有特别之处，略
//位运算符（ & | ^ << >> ） 在都转化为二进制的情况下有以下规则
// & 都是1时取1
// | 有1时取1
// ^ 不同时取1
// << 左移n位就是乘以2的n次方
// >> 右移n位就是除以2的n次方
//赋值运算符（ = += -= *= /= %= <<= >>= &= ^= |= ）
//其他运算符（ &返回变量存储地址 *指针变量 ）
// &+变量名 用来获取内存地址，使用 *+变量名 定义的对象来接收（此时用*运算接收后的对象，会变成变量的值，而不再是一个指针）
var ptr *bool

//func main() {
//	//newObj()
//	//editObj()
//	logic1()
//}

//对象的声明
func newObj() {
	//基本数据类型
	fmt.Println("hello")
	fmt.Println(bool1)
	fmt.Println(bool2, bool3)
	fmt.Println(int1)
	fmt.Println(uint1)
	fmt.Println(float321)
	fmt.Println(float641)
	fmt.Println(string1)
	//复杂数据类型
	fmt.Println(int_)
	fmt.Println(intArray)
	fmt.Println(mapLength)
	fmt.Println(intChan)
	fmt.Println(int_)
	fmt.Println(intFunc)
	fmt.Println(errorRealize)

	//其它声明方式
	fmt.Println(int2, int3)
	bool4 := true
	fmt.Println(bool4)
	//这种不带声明格式的只能在函数体中出现
	bool5, string2 := true, "hhhh"
	fmt.Println(bool5, string2)

	//常量值
	fmt.Println(enum1)
	fmt.Println(enum2)
	fmt.Println(enum3)
	fmt.Println(enum4)
	fmt.Println(enum5)
	fmt.Println(&bool5)
	ptr = &bool5
	fmt.Println(ptr)
	var ptr1 bool = *ptr
	fmt.Println(ptr1)
}

//对象赋值
func editObj() {
	editInt1 := 1
	editInt2 := editInt1
	editInt2 = 2
	fmt.Println(editInt1)
	fmt.Println(editInt2)
	println(editInt2)
	//交换变量
	editInt1, editInt2 = editInt2, editInt1
	fmt.Println(editInt1)
	fmt.Println(editInt2)
	//报错，值未声明
	//editInt3,editInt4 := editInt4,editInt3
	//"_"用于抛弃值 如方法返回值多个，不需要接收时
	_, editInt3 := 5, 7
	fmt.Println(editInt3)
	editInt1, editInt2, _ = func1() //只获取函数返回值的后两个
	fmt.Println(editInt1, editInt2)
}

func func1() (int, int, int) {
	a, b, c := 1, 2, 3
	return a, b, c
}

//一些逻辑运算
func logic1() {
	if !bool1 {
		fmt.Println(bool1)
	} else {
		if bool2 {
			fmt.Println(bool1)
		}
	}
	//golang中switch自带break 使用fallthrough关键字才能继续执行下面的case，这一点与java相反
	switch {
	case !(int1 < 0):
		fmt.Println("case !(int1 < 0)")
		fallthrough
	case int1 == 0:
		fmt.Println("case int1 == 0")
	default:
	}
	//用switch判断类型
	var x interface{}
	switch i := x.(type) {
	case nil:
		fmt.Println(" x 的类型 :%T", i)
	case int:
		fmt.Println("x 是 int 型")
	case float64:
		fmt.Println("x 是 float64 型")
	case func(int) float64:
		fmt.Println("x 是 func(int) 型")
	case bool, string:
		fmt.Println("x 是 bool 或 string 型")
	default:
		fmt.Println("未知型")
	}

	//var cc1 int
	//var chanInt1 chan int
	//select {
	//	case cc1= (<- chanInt1):
	//		fmt.Println()
	//}

	//select 语句类似于 switch 语句，但是select会随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Println("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Println("sent ", i2, " to c2\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Println("received ", i3, " from c3\n")
		} else {
			fmt.Println("c3 is closed\n")
		}
	default:
		fmt.Println("no communication\n")
	}

	//goto关键字 跳转到指定地方
	for int1 < 2 {
		fmt.Println(int1)
		int1++
		goto start
	}
start:
}
