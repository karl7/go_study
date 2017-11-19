package main

import (
	"fmt"
)

// 类型的别名
type (
	byte int8
	rune int32
	文本 string
)

// 对象提供的函数，大写字母开头
func Sum(i1 int, i2 int) int {
	return i1 + i2
}
// 函数可以有多个返回值
func sayHello(s1 string, s2 string)(s3 string, s4 string) {
	return s1, s2
}

func main() {
	// 变量声明方式1
	var a int = 1
	fmt.Println(a)
	// 变量声明方式2
	var a2 = 2
	fmt.Println(a2)
	// 变量声明方式3 最简洁方式
	a3 := 3
	fmt.Println(a3)

	// 使用自定义类型别名声明变量
	var b 文本
	b = "这是中文"
	fmt.Println("文本b：" + b)

	// 调用单个返回值的自定义函数
	fmt.Println(Sum(1, 2))

	// 调用多个返回值的自定义函数
	hello1, world2 := sayHello("hello", "world")
	fmt.Println(hello1 + world2)
	// return hello world
	fmt.Println(sayHello("hello", "world"))

	// 类型转换,return A
	fmt.Println(string(65))
	fmt.Println(int('A'))
}