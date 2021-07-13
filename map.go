package main

import "fmt"

func mainq() {

	map1:=make(map[int]string)
	map1[1]="java"
	map1[2]="c"
	map1[3]="php"

	fmt.Println(map1)

	map3:=map[int]string{
		1:"java",
		2:"gg",
		3:"ppp",
	}
	fmt.Println(map3)
}