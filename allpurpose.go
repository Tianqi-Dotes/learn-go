package main

import "fmt"

func allargs(arg interface{})  {

	fmt.Println(arg)

	_,isString:= arg.(string)

	if isString{
		fmt.Println("arg is string type")
	}else {
		fmt.Println("arg is not string type and type is %T",arg)
	}


}
func mainww() {
	allargs(1)
	allargs("arg")
}
