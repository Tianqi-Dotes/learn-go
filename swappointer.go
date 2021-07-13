package main

import "fmt"

func main6() {

	var a,b=10,20
	swap(&a,&b)
	fmt.Println(a,b)
}
func swap(aa *int, bb *int) {
	var temp int;
	temp=*aa
	*aa=*bb
	*bb=temp
}
