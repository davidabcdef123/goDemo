package main

import "fmt"

func main() {
	var b int =15
	var a int
	numbers:=[6]int{1,2,3,5}

	for a:=0;a<10;a++{
		fmt.Printf("第一个循环a的值是：%d\n",a)
	}

	for a<b{
		a++
		fmt.Printf("第二个循环a的值为：%d\n",a)
	}

	for i,x:=range numbers{
		fmt.Printf("第三次循环第 %d位 x的值=%d\n",i,x)
	}
	fmt.Println("----------------------------")
	fmt.Println(a,b,numbers)

}
