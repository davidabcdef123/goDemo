package main

import "fmt"

func main() {
	//截取
	//qiePianJieQu()
	appendNumbers()
}

func qiePianJieQu()  {
	numbers:=[]int{0,1,2,3,4,5,6,7,8,9}
	printSlice2(numbers)
	//原始
	fmt.Println("numbers==",numbers)
	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4]==",numbers[1:4])
	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])
	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1:=make([]int,0,5)
	printSlice2(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	numbers2:=numbers[:2]
	printSlice2(numbers2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice2(number3)

}

func appendNumbers(){
	var numbers []int
	printSlice2(numbers)

	numbers=append(numbers,0)
	printSlice2(numbers)
	numbers=append(numbers,1)
	printSlice2(numbers)
	/* 同时添加多个元素 */
	numbers = append(numbers, 2,3,4)
	printSlice2(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1:=make([]int,len(numbers),(cap(numbers))*2)
	copy(numbers1,numbers)
	printSlice2(numbers1)
}


func printSlice2(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v \n", len(x), cap(x), x)
}