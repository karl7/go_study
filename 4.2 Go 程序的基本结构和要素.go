package main

// https://github.com/karl7/the-way-to-go_ZH_CN/blob/master/eBook/04.2.md

// fm是包的别名
import fm "fmt"

// 多种包导入方式
// 1. 导入多个包
//import "fmt"
//import "os"
// 或者
//import "fmt";import "os"

// 2. 更短且更优雅的方式
//import (
//	"fmt"
//	"os"
//)
// 使用 gofmt 后将会被强制换行成以上格式
//import ("fmt"; "os")

// 如果包名不是以 . 或 / 开头，如 "fmt" 或者 "container/list"，则 Go 会在全局文件进行查找；如果包名以 ./ 开头，则 Go 会在相对目录中查找；如果包名以 / 开头（在 Windows 下也可以这样使用），则会在系统的绝对路径中查找。

func main() {
	fm.Println("hello world!")
}

// 可见性规则，大写开头代表public(其他包引用可见)，小写开头代表private(仅本包可用)
func Sum(a, b int) int { // 左大括号 { 必须与方法的声明放在同一行，这是编译器的强制规定
	return a + b
}

func Sum2(a, b int) (int, int, int) { // 左大括号 { 必须与方法的声明放在同一行，这是编译器的强制规定
	return a, b, a + b
}