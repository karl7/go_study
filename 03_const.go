package main

import "fmt"

const Monday, Tuesday, Wednesday, Thursday, Friday, Saturday = 1, 2, 3, 4, 5, 6

const (
	// iota 可被用作枚举，会在每次调用时自动累加1，但是跨const时，会从0重新计数
	Monday1 = iota // 0
	Tuesday1 // 1
	Wednesday1 // 2
	Thursday1 // 3
	Friday1 // 4
	Saturday1 = iota // 5
)

// 一个中文字符的len=3
const (str1 = "你好"; strSize = len(str1))

func main() {
	fmt.Println(Monday1)
	fmt.Println(Wednesday1)
	fmt.Println(Saturday1)
	// return 6
	fmt.Println(strSize)
}