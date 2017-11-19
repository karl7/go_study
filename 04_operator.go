package main

import "fmt"

/*
一元运算符： ^ !
二元运算符： << >> + - * / % | ^ & &^ == != >= <= < > && ||
channle： <-
 */
func main() {
	// 取反
	fmt.Println(!true)
	//
	fmt.Println(^2)
	fmt.Println(1^2)
	// 移位运算符，return 1024
	fmt.Println(1 << 10)
	/*
	 6: 0110
	11: 1011
	--------
	 &: 0010	and与，相同位的两个数字都为1，则为1；若有一个不为1，则为0
	 |: 1111	or或，相同位只要一个为1即为1
	 ^: 1101	xor异或，相同位不同则为1，相同则为0
	&^: 0100    and not与非，X and not Y等价于X and (not Y)
	 */
	 fmt.Println(6 & 11)
	fmt.Println(6 | 11)
	fmt.Println(6 ^ 11)
	// 以下2者等价
	fmt.Println(6 & (^11))
	fmt.Println(6 &^ 11)
}
