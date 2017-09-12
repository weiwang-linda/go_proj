package main

import (          //The "()" is must
	"fmt"
	"strconv"  //package for convert data type
)

func main() {

	/* declaration variable */
	var a int
	var b bool

	/* value to variable */
	a = 15
	b = false

	fmt.Println(strconv.Itoa(a))       //Println include "\n"
	fmt.Println(strconv.FormatBool(b))
}
