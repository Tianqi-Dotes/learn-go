package main

import "fmt"

type Human struct {
	name string
	age int
}

func (this *Human) Eat()  {

	fmt.Println("human eating~~")
}
func (this *Human) Run()  {

	fmt.Println("human running~~")
}

type SuperMan struct {
	Human
	lvl int
}

func (this *SuperMan) Eat() {
	fmt.Println("superman eatting")
}
func (this *SuperMan) Fly() {
	fmt.Println("superman flying")
}

func mainee() {
	mortal :=Human{"tq",27}
	mortal.Eat()
	mortal.Run()

	tq:=SuperMan{mortal,20}
	tq.Eat()
	tq.Fly()
}



