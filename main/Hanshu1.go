package main

import "fmt"

func max(num1 ,num2 int)int{
	var result int
	if num1>num2{
		result=num1
	}else if (num2>num1){
		result=num2
	}
	return result

}

func main(){
	var a=100
	var b=200
	var ret int
	ret=max(a,b)
	fmt.Println(ret)
}
