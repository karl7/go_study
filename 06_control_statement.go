package main

import "fmt"

func main() {
	a := 1
	for {
		a++
		fmt.Printf("%d %s", a, " ")
		if a > 10 {
			break
		}
	}
	fmt.Println()
	for i := 1; i < 10; i++ {
		fmt.Printf("%d %s", i, " ")
	}
	fmt.Println()

	// switch 使用形式之一
	b := 1
	switch b {
	case 0:
		fmt.Println(0)
	case 1:
		fmt.Println(1)
		fallthrough // 表示继续执行后续case中操作
	case 2:
		fmt.Println(2)
	default:
		fmt.Println(3)
	}
	// switch 使用形式之二
	switch {
	case a>= 0:
		fmt.Println(222)
	default:
		fmt.Println(111)
	}

	// 标签的使用，如果把continue换成goto结果会有什么不一样？
	LABEL:
		for j := 0; j < 10; j++ {
			for {
				fmt.Println(j)
				continue LABEL
			}
		}
}