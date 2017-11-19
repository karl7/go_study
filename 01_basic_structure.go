// 当前程序的包名
package main

// 导入其他的包
import "fmt"
import (
	"os"
	"runtime"
)

// 常量的定义
const PI = 3.14159
// 常量组
const (
	PI2 = 3.1415926
	PI3 = 3.141592678
)

// 全局变量的声明与赋值
var name = "gopher"
// 全局变量组
var (
	name2 = "golang"
	name3 = "goland"
)

// 一般类型声明
type newType int16
type (
	newType2 int32
	newType3 int64
)

// 结构的声明
type gopher struct{}

// 接口的声明
type golang interface{}

func main() {
	fmt.Println("Hello world!")

	var goos string = runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
}
