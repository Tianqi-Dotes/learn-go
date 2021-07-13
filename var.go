package main

import "fmt"

var g="gg"
var(
	gg="gggg"
	qq="qq"
)

func main2()  {
	var a int;
	fmt.Println("a= ",a)
	fmt.Printf("a type is %T \n",a)

	var b int=100;
	fmt.Println("b= ",b)
	fmt.Printf("b type is %T \n",b)

	var c =100;
	fmt.Println("b= ",c)
	fmt.Printf("c type is %T \n",c)

	var s string="abcd"
	fmt.Printf("s type is %T \n",s)

	ss:="qwerty"
	fmt.Printf("ss type is %T \n",ss)

	var kk,yy=11,22;
	fmt.Println("kk value is ",kk,"yy value is ", yy)

	var cc,vv=11,"tq";
	fmt.Println("cc value is ",cc,"vv value is ", vv)


	fmt.Println(g)
	fmt.Println(gg)
	fmt.Println(qq)
}
