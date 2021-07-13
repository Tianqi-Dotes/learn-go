package main

import (
	"fmt"
	"time"
)

func mainzc() {


	c:=make(chan int,3)

	go func() {
		defer fmt.Println("child process end")

		for i:=0;i<3;i++{
			c<-i
			fmt.Println("child process send data,now the length is ",len(c)," and the cap is ",cap(c))
		}
	}()

	time.Sleep(2*time.Second)

	for i:=0;i<3;i++{
		num:=<-c
		fmt.Println("main get data : ",num)
	}
}
