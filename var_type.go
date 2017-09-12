package main

import (          //The "()" is must
	"fmt"
	"strconv"  //package for convert data type
)

func main() {

	/* declaration variable */
	var a int
	var b bool

	/* assign value to variable */
	a = 15
	b = false
	f := 3.14      //declaration and assign value in one line.
	var (
		x int
		y bool
	)                //like import and const

	m, n := 20, 16    //assign two variables in one time.
	_, t := 34, 35    //_ is special varialbe, any value assigned to it will be lost.

	fmt.Println(strconv.Itoa(a))       //Println include "\n"
	fmt.Println(strconv.FormatBool(b))   //convert bool to strong
	fmt.Println(strconv.FormatFloat(f, 'E', -1, 64))    //convert float to string, ParseFloat(f,64) is convert string to float
}
