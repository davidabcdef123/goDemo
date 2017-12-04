package main

import (
	"fmt"
	"unsafe"
)

const (
	ac="abc"
	bc=len(ac)
	cc=unsafe.Sizeof(ac)
)

const (
	ad=iota
	bd=iota
	cd=iota
)
//等价于
const (
	ae=iota
	be
	ce
)

func main(){
	const  LENGTH int  =10
	const  WIDTH int  =5
	var area int
	const a,b,c  = 1,false,"str"
	area=LENGTH*WIDTH
	fmt.Printf("面积为 : %d", area)
	fmt.Printf("面积为:%d",area)
	fmt.Println()
	fmt.Println(ac,bc,cc)
}

