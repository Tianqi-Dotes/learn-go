package main

import "fmt"

type Hero struct {
	name string
	lvl int
	att int
}

func (this *Hero) info(){
	fmt.Println(this)
}

func (this *Hero) changeName(newname string) string{
	this.name=newname
	return newname
}
func (this *Hero) addAtt(attt int){
	this.att+=attt
}

func maina() {
	tq :=Hero{"tq",15,99}
	tq.info()

	tq.addAtt(1)
	tq.info()

	tq.changeName("tqq")
	tq.info()
}
