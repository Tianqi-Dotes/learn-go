package main

import "fmt"

func addElement(mmap map[string]string)  {
	mmap["gg"]="gg"
}

func mainw() {
	cityMap :=make(map[string]string)

	cityMap["china"]="beijing"
	cityMap["usa"]="washinton"
	cityMap["uk"]="london"

	for k,v :=range cityMap{
		fmt.Println(k,v)
	}

	fmt.Println("-------------------")
	delete(cityMap,"usa")
	for k,v :=range cityMap{
		fmt.Println(k,v)
	}
	fmt.Println("-------------------")
	cityMap["usa"]="newyork"
	for k,v :=range cityMap{
		fmt.Println(k,v)
	}
	fmt.Println("-------------------")
	addElement(cityMap)
	for k,v :=range cityMap{
		fmt.Println(k,v)
	}


}
