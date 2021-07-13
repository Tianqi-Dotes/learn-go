package main

import "fmt"

func main5() {

	var a=1;
	changeValue(&a)

	fmt.Println(a)

}
func changeValue(p *int) {
	*p=10
}
