package main

import "fmt"

func main10() {
	s:=[]int{1,2,3,4}
	s1:=s[0:2]

	fmt.Println(s1)
	s1[0]=100

	fmt.Println(s)

	s2:=make([]int,4,10)
	fmt.Printf("len is %d, capacity is %d, values are %v  \n",len(s2),cap(s2),s2)
	copy(s2,s)
	fmt.Println(s2)


}
