package main

import "fmt"

var a int = 10
var b = 10

var a1="aaaa"
var b1 string="aaaabbb"
var c bool
//cx123 : = 10

var x,y int
var (// 这种因式分解关键字的写法一般用于声明全局变量
		a2 int
		b2 bool
)
var c1,d1 int =1,2
var e,f=123,"helloworld"
//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"


func main() {
	fmt.Println("aaaa")
	fmt.Println(a1,b1,c)

}

func init() {
	fmt.Print("init()")
	test2()
	//j, k := 123, "hello"

}

func test2() {
	fmt.Print("test2")
	//ga:=123
}

//如果你声明了一个局部变量却没有在相同的代码块中使用它，同样会得到编译错误
func test3(){
	//var s string= "abc"
	//fmt.Println(s)
}
