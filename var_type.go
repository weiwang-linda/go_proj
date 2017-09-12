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

	fmt.Println(strconv.Itoa(a))       //Println include "\n"
	fmt.Println(strconv.FormatBool(b))   //convert bool to strong
	fmt.Println(strconv.FormatFloat(f, 'E', -1, 64))    //convert float to string, ParseFloat(f,64) is convert string to float
}
