package main

import "fmt"

func printArray(arr1 []int)  {

	for _,value := range arr1{

		fmt.Println("value is ",value)

	}

	arr1[0]=100
}

func main0() {

	arr := []int{1,2,3,4}
	printArray(arr)

	for k,v := range arr{
		fmt.Println("key: ",k," value ",v)
	}


	slice1 :=make([]int,3)
	fmt.Printf("slice is length %d and values are %v",len(slice1),slice1)


}