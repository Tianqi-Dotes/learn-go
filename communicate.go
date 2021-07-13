package main

import "fmt"

func mainzzz() {
	c:=make(chan int)

	go func() {
		defer fmt.Println("child goroutine finished")

		fmt.Println("child goroutine running")
		c<-666

	}()

	num :=<-c
	fmt.Println("received data: ",num)

	fmt.Println("main finished")

}
