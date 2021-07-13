package main

import "fmt"

func foo1(a string, b int) int {

	fmt.Println(a)
	fmt.Println(b)
	c:=100
	return c
}

func foo2(a string, b int) (c int, d int) {
	fmt.Println(a)
	fmt.Println(b)
	c=1000
	d=2000
	return
}
func main4() {
	c:=foo1("s",2);
	fmt.Println(c)

	c,d:=foo2("s",1)
	fmt.Println(c,d)
}
