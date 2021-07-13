package main

import "fmt"

const (
	QQ=iota
	PP
	ZZ
	MM
)
const(
	a,b=iota+1,iota+2
	c,d
	e,f
	ggg,h

	i,j=iota*2,iota*3
	k,l
	m,n
)
func main3() {
	const length int = 10;

	fmt.Println(length);
	fmt.Println(QQ)
	fmt.Println(PP)
	fmt.Println(ZZ)
	fmt.Println(MM)

	fmt.Println(a,b)
	fmt.Println(c,d)
	fmt.Println(e,f)
	fmt.Println(ggg,h)

	fmt.Println(i,j)
	fmt.Println(k,l)
	fmt.Println(m,n)

}
