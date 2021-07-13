package main

import "fmt"

func main7()  {
	var array1 [10]int;
	array2:=[5]int{1,2,3,4}

	for i:=0;i<len(array1);i++{
		fmt.Println("index ",i," value: ",array1[i])
	}

	for index,v:=range array2{
		fmt.Println("index ",index," value: ",v)
	}
}
