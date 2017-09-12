/* constant is created at build time.
   Only can be number, string and bool.
   Using iota to create emenu.
   iota == the smallest, [iota] == "not a iota" */

package main

import (
	"fmt"
	"strconv"
)

func main(){
	const (
		a = iota      //iota is the smallest number
		b = iota      //the second time using iota, the number will plug 1
	)

	const (
		x = iota
		y             //omit '= iota'
	)

	const (               //fixed, and specify type for constant
		m = 0            //implicit
		n string = "0"   //explicit
	)

	fmt.Println(strconv.Itoa(a))
	fmt.Println(strconv.Itoa(b))
	fmt.Println(strconv.Itoa(x))
	fmt.Println(strconv.Itoa(y))
	fmt.Println(strconv.Itoa(m))
	fmt.Println(n)
}
