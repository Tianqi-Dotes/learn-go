package main

import "fmt"

type Book struct {
	author string
	pages int
}

func changeAuthor(b *Book)  {
	b.author="tqq"
}

func maine() {

	b := Book{"tq",100}
	fmt.Println(b)

	changeAuthor(&b)
	fmt.Println(b)
}
