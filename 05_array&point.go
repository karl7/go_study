package main

import "fmt"

func main() {
	a := 1
	// &a 取地址操作
	fmt.Println(&a)
	// 指针变量
	var p *int = &a
	fmt.Println(p)

	// 数组指针
	arr1 := [...]int{1,2,3}
	// var arr1point *[3]int = &arr1
	arr1point := &arr1
	fmt.Println(arr1point)

	// 指针数组
	b, c := 3, 4
	// var arr2point [2]*int = [...]*int{&b, &c}
	arr2point := [...]*int{&b, &c}
	fmt.Println(arr2point)
}