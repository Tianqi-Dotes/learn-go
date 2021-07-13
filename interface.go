package main

import "fmt"

type animal interface {
	bark()
	getcolor()
	getType()
	getInfo()
}

type Dog struct {
	t string
	color string
}
type Cat struct {
	t string
	color string
}

func (this *Dog)bark () {
	fmt.Println("wang wang")
}
func (this *Cat)bark()  {
	fmt.Println("miao")
}
func (this *Cat)getcolor()  {
	fmt.Println("color is ",this.color)
}
func (this *Dog)getcolor()  {
	fmt.Println("color is ",this.color)
}
func (this *Dog)getType()  {
	fmt.Println("animal is ",this.t)
}
func (this *Cat)getType()  {
	fmt.Println("animal is ",this.t)
}
func (this *Dog)getInfo()  {
	fmt.Println(this)
}
func (this *Cat)getInfo()  {
	fmt.Println(this)
}

func mainf() {
	var animal animal
	animal=&Dog{"dog","yellow"}
	animal.getInfo()
	animal.bark()
	animal=&Cat{"cat","grey"}
	animal.getInfo()
	animal.bark()

}

