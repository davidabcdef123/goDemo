package main

import "fmt"

func main() {

	var countryCaptalMap map[string]string

	//创建集合
	countryCaptalMap=make(map[string]string)

	countryCaptalMap["France"]="Paris"
	countryCaptalMap["Italy"]="Rome"
	countryCaptalMap["Japan"]="Tokyo"
	countryCaptalMap["India"]="New Delhi"

	for country :=range countryCaptalMap{
		fmt.Println(country,"国家的首都是",countryCaptalMap[country])
	}

	for a,b:=range countryCaptalMap{
		fmt.Println(a,b)
	}

	/* 查看元素在集合中是否存在 */
	captial,aaa:=countryCaptalMap["India"]
	if aaa{
		fmt.Println("有首都是",captial)
	}else{
		fmt.Println("没有")
	}
	deleteCountry()



}

func deleteCountry(){

	countryCapitaiMap:=map[string]string{"France":"Paris","Italy":"Rome","Japan":"Tokyo","India":"New Delhi"}
	fmt.Println("原始map")
	for country:=range countryCapitaiMap{
		fmt.Println(country,"首都是",countryCapitaiMap[country])
	}

	delete(countryCapitaiMap,"France")
	fmt.Println("Entry for France is deleted")
	fmt.Println("删除后元素")
	for country := range countryCapitaiMap{
		fmt.Println(country,"首都是",countryCapitaiMap[country])

	}

}