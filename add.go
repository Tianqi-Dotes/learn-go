package main

import (
	"fmt"
)

func add(a,b chan int)  {
	x,y:=1,1

	for  {
		select {//select 监控多个channel
		case a <- x:
			x=y
			y=x+y
		case <-b:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c:=make(chan int)
	quit:=make(chan int)

	go func() {
		for i:=0;i<10;i++ {

			fmt.Println(<-c)

		}
		quit<-0
	}()

	add(c,quit)

}
